package main

import (
  "errors"
  "fmt"
)

// Sentinel errors
var ErrNotFound = errors.New("not found")
var ErrPermission = errors.New("permission denied")

// fetchData simulates fetching data with possible errors
func fetchData(resource string) error {
  switch resource {
  case "secret":
    return fmt.Errorf("access to %s: %w", resource, ErrPermission)
  case "missing":
    return fmt.Errorf("resource %s: %w", resource, ErrNotFound)
  default:
    return nil
  }
}

func main() {
  fmt.Println("=== errors.Is - Check Error Type ===")

  // Check for specific error using errors.Is
  err := fetchData("missing")
  if errors.Is(err, ErrNotFound) {
    fmt.Println("Resource was not found")
  }

  err = fetchData("secret")
  if errors.Is(err, ErrPermission) {
    fmt.Println("Permission was denied")
  }

  // errors.Is works through wrapped errors
  fmt.Println("\n=== Wrapped Error Check ===")
  err = fetchData("missing")
  fmt.Println("Error:", err)
  fmt.Println("Is ErrNotFound?", errors.Is(err, ErrNotFound))
  fmt.Println("Is ErrPermission?", errors.Is(err, ErrPermission))

  fmt.Println("\n=== Handle by Error Type ===")

  resources := []string{"data", "missing", "secret"}
  for _, r := range resources {
    err := fetchData(r)
    switch {
    case err == nil:
      fmt.Printf("%s: fetched successfully\n", r)
    case errors.Is(err, ErrNotFound):
      fmt.Printf("%s: not found, try later\n", r)
    case errors.Is(err, ErrPermission):
      fmt.Printf("%s: need higher access\n", r)
    default:
      fmt.Printf("%s: unknown error: %v\n", r, err)
    }
  }
}
