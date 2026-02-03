package main

import (
  "errors"
  "fmt"
)

// Sentinel error - a predefined error value
var ErrNotFound = errors.New("item not found")

// findUser simulates finding a user by ID
func findUser(id int) (string, error) {
  users := map[int]string{
    1: "Alice",
    2: "Bob",
    3: "Carol",
  }

  name, exists := users[id]
  if !exists {
    // Wrap the sentinel error with context
    return "", fmt.Errorf("user ID %d: %w", id, ErrNotFound)
  }
  return name, nil
}

// getProfile wraps the findUser error with more context
func getProfile(id int) (string, error) {
  user, err := findUser(id)
  if err != nil {
    return "", fmt.Errorf("getProfile failed: %w", err)
  }
  return fmt.Sprintf("Profile: %s", user), nil
}

func main() {
  fmt.Println("=== Error Wrapping ===")

  // Successful lookup
  profile, err := getProfile(1)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println(profile)
  }

  // Failed lookup
  profile, err = getProfile(99)
  if err != nil {
    fmt.Println("Error:", err)
  }

  fmt.Println("\n=== Error Chain ===")

  // The wrapped error preserves the full chain
  _, err = getProfile(99)
  if err != nil {
    fmt.Println("Full error:", err)
    fmt.Println("Unwrapped:", errors.Unwrap(err))
  }
}
