package main

import (
  "errors"
  "fmt"
)

// divide returns an error if dividing by zero
func divide(a, b float64) (float64, error) {
  if b == 0 {
    return 0, errors.New("cannot divide by zero")
  }
  return a / b, nil
}

// validateAge returns an error if age is invalid
func validateAge(age int) error {
  if age < 0 {
    return errors.New("age cannot be negative")
  }
  if age > 150 {
    return errors.New("age seems unrealistic")
  }
  return nil
}

func main() {
  fmt.Println("=== Error Basics ===")

  // Successful division
  result, err := divide(10, 2)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("10 / 2 =", result)
  }

  // Division by zero
  result, err = divide(10, 0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("10 / 0 =", result)
  }

  fmt.Println("\n=== Validate Age ===")

  // Valid age
  err = validateAge(25)
  if err != nil {
    fmt.Println("Invalid:", err)
  } else {
    fmt.Println("Age 25 is valid")
  }

  // Negative age
  err = validateAge(-5)
  if err != nil {
    fmt.Println("Invalid:", err)
  } else {
    fmt.Println("Age -5 is valid")
  }

  // Unrealistic age
  err = validateAge(200)
  if err != nil {
    fmt.Println("Invalid:", err)
  } else {
    fmt.Println("Age 200 is valid")
  }
}
