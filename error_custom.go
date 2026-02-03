package main

import (
  "errors"
  "fmt"
)

// ValidationError is a custom error type
type ValidationError struct {
  Field   string
  Message string
}

// Error implements the error interface
func (e ValidationError) Error() string {
  return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// validateUser returns a ValidationError for invalid input
func validateUser(name string, age int) error {
  if name == "" {
    return ValidationError{Field: "name", Message: "cannot be empty"}
  }
  if age < 0 {
    return ValidationError{Field: "age", Message: "cannot be negative"}
  }
  if age > 150 {
    return ValidationError{Field: "age", Message: "must be realistic"}
  }
  return nil
}

func main() {
  fmt.Println("=== Custom Error Types ===")

  // Valid user
  err := validateUser("Alice", 25)
  if err != nil {
    fmt.Println("Validation failed:", err)
  } else {
    fmt.Println("User is valid")
  }

  // Invalid name
  err = validateUser("", 25)
  if err != nil {
    fmt.Println("Validation failed:", err)
  }

  // Invalid age
  err = validateUser("Bob", -5)
  if err != nil {
    fmt.Println("Validation failed:", err)

    // Use errors.As to get the underlying type
    var valErr ValidationError
    if errors.As(err, &valErr) {
      fmt.Println("  Field:", valErr.Field)
      fmt.Println("  Message:", valErr.Message)
    }
  }

  fmt.Println("\n=== errors.As Example ===")

  err = validateUser("Carol", 200)
  var valErr ValidationError
  if errors.As(err, &valErr) {
    fmt.Printf("Field %q failed: %s\n", valErr.Field, valErr.Message)
  }
}
