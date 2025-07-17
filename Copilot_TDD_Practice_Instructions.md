# Copilot_TDD_Practice_Instructions.md

# GitHub Copilot TDD Practice Instructions for Go User Profile REST API

Follow these tasks to practice Test-Driven Development (TDD) using GitHub Copilot's custom chat modes within your Go REST API project.

---

## ✅ Task 1: Implement Unit Tests Using Testing Custom Chat Mode

**Goal:** Create comprehensive unit tests for the existing User Profile API endpoints using Copilot's Testing custom chat mode.

### Step-by-Step Instructions:

1. **Open Copilot Chat Panel**
   - Open the Copilot Chat panel in VS Code (Ctrl+Shift+I or Cmd+Shift+I)
   - Click on the chat mode selector at the top of the chat panel

2. **Select Testing Custom Chat Mode**
   - Choose the "Testing" custom chat mode from the dropdown
   - This mode is optimized for generating test code and testing strategies

3. **Generate Unit Tests with Simple Prompt**
   - Use this prompt in the Testing chat mode:

   ```
   Generate and run unit tests for me
   ```

4. **Review and Accept Generated Tests**
   - Copilot will generate a comprehensive test file
   - Review the generated code for completeness
   - Accept the suggested test file creation

5. **Run the Tests**
   - Open the integrated terminal
   - Navigate to the project root
   - Run the tests with: `go test ./controllers -v`
   - Verify all tests pass

6. **Analyze Test Coverage**
   - Run tests with coverage: `go test ./controllers -cover`
   - Aim for >80% test coverage

---

## ✅ Task 2: Implement Delete Functionality Using TDD Custom Chat Mode

**Goal:** Use Test-Driven Development approach to implement the missing `DeleteUser` functionality by writing tests first, then implementing the feature.

### Step-by-Step Instructions:

> **Note:** Copilot may combine several of these steps and complete them without you specifically prompting for it. Keep an eye on what is happening and the end result should have all of these steps completed.

1. **Switch to TDD Custom Chat Mode**
   - In the Copilot Chat panel, change the chat mode to "TDD"
   - This mode follows Test-Driven Development principles

2. **Request Delete Functionality Implementation**
   - Provide this prompt to the TDD chat mode:

   ```
   Implement a DeleteUser function for the REST API
   ```

   - The TDD chat mode will automatically intervene and suggest writing tests first
   - Follow Copilot's guidance to create tests before implementation

3. **Review and Accept Test Cases (Red Phase)**
   - Accept the generated test cases that Copilot suggests
   - Ensure tests cover all scenarios:
     - Successful deletion
     - User not found
     - Edge cases

4. **Run Tests to Confirm They Fail (Red Phase)**
   - Run the tests: `go test ./controllers -v`
   - Confirm the DeleteUser tests fail (as expected, since function isn't implemented)
   - This validates we're in the "Red" phase of TDD

5. **Proceed with Implementation (Green Phase)**
   - Once tests are failing, ask Copilot to continue with the implementation:

   ```
   Now implement the DeleteUser function to make the tests pass
   ```

6. **Add Route Configuration**
   - Ask Copilot to update the routes:

   ```
   Add the DELETE route for the DeleteUser function
   ```

7. **Run Tests to Confirm They Pass (Green Phase)**
   - Run the tests again: `go test ./controllers -v`
   - Verify all tests now pass
   - Confirm we've reached the "Green" phase

8. **Refactor if Needed (Refactor Phase)**
   - Ask Copilot to review the implementation:

   ```
   Review the DeleteUser implementation for potential improvements
   ```

9. **Integration Testing**
   - Start the server: `go run main.go`
   - Test the delete endpoint manually:
     ```bash
     # Create a user first
     curl -X POST http://localhost:8080/api/v1/users \
       -H "Content-Type: application/json" \
       -d '{"id":"test123","fullName":"Test User","emoji":"🧪"}'
     
     # Delete the user
     curl -X DELETE http://localhost:8080/api/v1/users/test123 -v
     
     # Verify user is deleted
     curl http://localhost:8080/api/v1/users/test123
     ```

10. **Final Test Run**
    - Run all tests one final time: `go test ./... -v`
    - Ensure 100% of tests pass
    - Check test coverage: `go test ./... -cover`

---

## 🎯 TDD Principles Reinforced

Through this exercise, you've practiced:

- **Red-Green-Refactor Cycle**: Writing failing tests first, implementing minimal code to pass, then refactoring
- **Test-First Development**: Defining behavior through tests before implementation
- **Incremental Development**: Building functionality step by step with immediate feedback
- **Custom Chat Modes**: Leveraging Copilot's specialized TDD and Testing modes for optimal code generation

---

## 🎉 Completion

You've successfully practiced Test-Driven Development using GitHub Copilot's custom chat modes! You now have:
- Comprehensive unit tests for your Go REST API
- A fully implemented delete functionality following TDD principles
- Experience with Copilot's Testing and TDD chat modes
- A robust testing foundation for future development
