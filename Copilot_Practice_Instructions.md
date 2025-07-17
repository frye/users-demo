# Copilot_Practice_Instructions.md

# GitHub Copilot Practice Instructions for Go User Profile REST API

Follow these tasks to practice using GitHub Copilot effectively within your Go REST API project.

---

## ✅ Task 1: Add Three More Sample Users (Inline Code Completion)

**File:** [`controllers/user_controller.go`](controllers/user_controller.go)

- Locate the existing `users` slice:

var users = []models.UserProfile{
    {ID: "1", FullName: "John Doe", Emoji: "😀"},
    {ID: "2", FullName: "Jane Smith", Emoji: "🚀"},
    {ID: "3", FullName: "Robert Johnson", Emoji: "🎸"},
    // Add three more users here using Copilot inline completion
}

- Use GitHub Copilot inline completion to add three additional unique users:
  - Trigger Copilot suggestions by starting a new line.
  - Accept or cycle through suggestions to complete the task.

---

## ✅ Task 2: Implement Delete Functionality (Ask or Edit Mode - Base Model)

**File:** [`controllers/user_controller.go`](controllers/user_controller.go)

- Locate the incomplete `DeleteUser` function:

// DeleteUser removes a user by ID TODO
func DeleteUser(c *gin.Context) {
    // Implement delete functionality here using Copilot Ask or Edit mode
}

- Use Copilot's Edit mode (base model):
 - Drag and drop routes.go and user_controller.go into the copilot chat window.
- Provide the following prompt to Copilot:
"Implement delete functionality to remove a user by ID from the users slice and return appropriate HTTP responses."

- Review and accept the generated code.

---

## ✅ Task 3: Use Copilot CLI to Delete a User (Integrated Terminal)

- Open the integrated terminal in Visual Studio Code.
- Use Copilot CLI to generate a curl command to delete a user via the API:

gh copilot suggest "Generate a curl command to delete user with ID '2' from the API at http://localhost:8080/api/v1/users/2"

- Execute the suggested curl command to verify the delete functionality.

---

## ✅ Task 4: Implement Unit Tests (Agent Mode - Premium Model)

- Activate Copilot Agent mode (use premium model like Claude 3.7 Sonnet for example) from the Copilot sidebar.
- Provide the following prompt to Copilot Agent:

"Generate and run unit tests for this application using Go's testing package and Gin's test utilities."

- Review, accept, and run the generated tests to ensure correctness.

---

## 🎉 Completion

You've successfully practiced using GitHub Copilot's inline completion, Ask/Edit mode, CLI suggestions, and Agent mode to enhance your Go REST API development workflow!