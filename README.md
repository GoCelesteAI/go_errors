# Go Error Handling Examples

Demo code from Go Tutorial Lesson 19: Error Handling

## Topics Covered

1. **error_basics.go** - Creating errors with `errors.New`, returning errors from functions
2. **error_wrap.go** - Error wrapping with `fmt.Errorf` and `%w`, sentinel errors
3. **error_check.go** - Checking errors with `errors.Is`, handling by error type
4. **error_custom.go** - Custom error types, `errors.As` for type extraction

## Running the Examples

```bash
go run error_basics.go
go run error_wrap.go
go run error_check.go
go run error_custom.go
```

## Key Concepts

- Go uses explicit error returns instead of exceptions
- The `error` interface requires only an `Error() string` method
- Use `errors.New()` for simple error messages
- Use `fmt.Errorf()` with `%w` to wrap errors with context
- Use `errors.Is()` to check if an error matches a sentinel
- Use `errors.As()` to extract custom error types

## YouTube Tutorial

Watch the full tutorial: [Go Lesson 19: Error Handling](https://www.youtube.com/@CelesteAI)

## License

MIT
