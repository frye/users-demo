package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"userprofile-api/models"
)

// setupTestRouter creates a test router with necessary configuration
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Setup template rendering for HomePageHandler test
	router.LoadHTMLFiles("../templates/users.html")
	
	return router
}

// resetUsers resets the users slice to original test data for test isolation
func resetUsers() {
	users = []models.UserProfile{
		{ID: "1", FullName: "John Doe", Emoji: "😀"},
		{ID: "2", FullName: "Jane Smith", Emoji: "🚀"},
		{ID: "3", FullName: "Robert Johnson", Emoji: "🎸"},
	}
}

func TestHomePageHandler(t *testing.T) {
	resetUsers()
	router := setupTestRouter()
	router.GET("/", HomePageHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Check that the response contains HTML content
	contentType := w.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected content type 'text/html; charset=utf-8', got '%s'", contentType)
	}
}

func TestGetUsers(t *testing.T) {
	resetUsers()
	router := setupTestRouter()
	router.GET("/api/v1/users", GetUsers)

	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var responseUsers []models.UserProfile
	err := json.Unmarshal(w.Body.Bytes(), &responseUsers)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(responseUsers) != 3 {
		t.Errorf("Expected 3 users, got %d", len(responseUsers))
	}

	// Verify first user data
	if responseUsers[0].ID != "1" || responseUsers[0].FullName != "John Doe" {
		t.Errorf("Unexpected user data: %+v", responseUsers[0])
	}
}

func TestGetUser(t *testing.T) {
	resetUsers()
	router := setupTestRouter()
	router.GET("/api/v1/users/:id", GetUser)

	tests := []struct {
		name           string
		userID         string
		expectedStatus int
		expectedUser   *models.UserProfile
		expectError    bool
	}{
		{
			name:           "Get existing user",
			userID:         "1",
			expectedStatus: http.StatusOK,
			expectedUser:   &models.UserProfile{ID: "1", FullName: "John Doe", Emoji: "😀"},
			expectError:    false,
		},
		{
			name:           "Get another existing user",
			userID:         "2",
			expectedStatus: http.StatusOK,
			expectedUser:   &models.UserProfile{ID: "2", FullName: "Jane Smith", Emoji: "🚀"},
			expectError:    false,
		},
		{
			name:           "Get non-existent user",
			userID:         "999",
			expectedStatus: http.StatusNotFound,
			expectedUser:   nil,
			expectError:    true,
		},
		{
			name:           "Get user with non-existent alphanumeric ID",
			userID:         "abc",
			expectedStatus: http.StatusNotFound,
			expectedUser:   nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/api/v1/users/"+tt.userID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectError {
				var errorResponse map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if errorResponse["error"] != "User not found" {
					t.Errorf("Expected error message 'User not found', got '%s'", errorResponse["error"])
				}
			} else {
				var responseUser models.UserProfile
				err := json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				if responseUser != *tt.expectedUser {
					t.Errorf("Expected user %+v, got %+v", *tt.expectedUser, responseUser)
				}
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	router := setupTestRouter()
	router.POST("/api/v1/users", CreateUser)

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectError    bool
	}{
		{
			name: "Create valid user",
			requestBody: models.UserProfile{
				ID:       "4",
				FullName: "Alice Cooper",
				Emoji:    "🎭",
			},
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name: "Create another valid user",
			requestBody: models.UserProfile{
				ID:       "5",
				FullName: "Bob Dylan",
				Emoji:    "🎵",
			},
			expectedStatus: http.StatusCreated,
			expectError:    false,
		},
		{
			name:           "Create user with invalid JSON",
			requestBody:    `{"invalid": json}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "Create user with empty body",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetUsers() // Reset for each test

			var reqBody []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, _ := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectError {
				var errorResponse map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if errorResponse["error"] == "" {
					t.Error("Expected error message, got empty")
				}
			} else {
				var responseUser models.UserProfile
				err := json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				
				expectedUser := tt.requestBody.(models.UserProfile)
				if responseUser != expectedUser {
					t.Errorf("Expected user %+v, got %+v", expectedUser, responseUser)
				}
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupTestRouter()
	router.PUT("/api/v1/users/:id", UpdateUser)

	tests := []struct {
		name           string
		userID         string
		requestBody    interface{}
		expectedStatus int
		expectError    bool
	}{
		{
			name:   "Update existing user",
			userID: "1",
			requestBody: models.UserProfile{
				FullName: "John Smith",
				Emoji:    "😎",
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:   "Update another existing user",
			userID: "2",
			requestBody: models.UserProfile{
				FullName: "Jane Doe",
				Emoji:    "🌟",
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
		{
			name:   "Update non-existent user",
			userID: "999",
			requestBody: models.UserProfile{
				FullName: "Ghost User",
				Emoji:    "👻",
			},
			expectedStatus: http.StatusNotFound,
			expectError:    true,
		},
		{
			name:           "Update user with invalid JSON",
			userID:         "1",
			requestBody:    `{"invalid": json}`,
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
		{
			name:           "Update user with empty body",
			userID:         "1",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetUsers() // Reset for each test

			var reqBody []byte
			var err error

			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				if err != nil {
					t.Fatalf("Failed to marshal request body: %v", err)
				}
			}

			req, _ := http.NewRequest("PUT", "/api/v1/users/"+tt.userID, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.expectError {
				var errorResponse map[string]string
				err := json.Unmarshal(w.Body.Bytes(), &errorResponse)
				if err != nil {
					t.Errorf("Failed to unmarshal error response: %v", err)
				}
				if errorResponse["error"] == "" {
					t.Error("Expected error message, got empty")
				}
			} else {
				var responseUser models.UserProfile
				err := json.Unmarshal(w.Body.Bytes(), &responseUser)
				if err != nil {
					t.Errorf("Failed to unmarshal response: %v", err)
				}
				
				// Verify the ID is set correctly (should be from URL, not request body)
				if responseUser.ID != tt.userID {
					t.Errorf("Expected user ID %s, got %s", tt.userID, responseUser.ID)
				}
				
				expectedUser := tt.requestBody.(models.UserProfile)
				if responseUser.FullName != expectedUser.FullName || responseUser.Emoji != expectedUser.Emoji {
					t.Errorf("Expected user data FullName=%s, Emoji=%s, got FullName=%s, Emoji=%s", 
						expectedUser.FullName, expectedUser.Emoji, responseUser.FullName, responseUser.Emoji)
				}
			}
		})
	}
}