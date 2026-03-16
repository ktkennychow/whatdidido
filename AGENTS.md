# AGENTS.md - Development Guidelines for whatdidido

## Build Commands

- `go build` - Build the binary
- `go build -o whatdidido` - Build with specific output name
- `go install github.com/ktkennychow/whatdidido@latest` - Install from remote

## Test Commands

- `go test` - Run all tests
- `go test -run TestFunctionName` - Run specific test
- `go test -v` - Run tests with verbose output
- `./test.sh` - Run comprehensive integration test script

## Lint Commands

- `go vet` - Check for common errors
- `go fmt` - Format code according to Go standards

## Code Style Guidelines

### General

- Follow standard Go formatting (`go fmt`)
- Use `gofmt` compatible formatting
- Keep lines under 100 characters when possible
- Use meaningful variable and function names

### Imports

- Standard library imports first
- Third-party imports second
- Local imports last
- Group imports by blank lines between groups

### Naming Conventions

- Variables and functions: camelCase
- Exported types/functions: PascalCase
- Constants: PascalCase
- Struct fields: PascalCase if exported, camelCase if unexported

### Error Handling

- Always check errors returned by functions
- Use `if err != nil` pattern immediately after function calls
- Return errors to caller rather than handling internally when appropriate
- Use `fmt.Errorf` for error wrapping

### Types and Structs

- Define config structs with JSON tags for serialization
- Use pointer receivers for methods that modify the receiver
- Keep struct definitions simple and focused

### Functions

- Keep functions focused on single responsibility
- Use descriptive names that indicate purpose
- Handle edge cases and invalid inputs gracefully

### Testing

- Write unit tests for core functions
- Test edge cases including empty inputs
- Use table-driven tests for multiple test cases
- Test configuration loading/saving thoroughly

### Dependencies

- Use cobra for CLI framework
- Keep dependencies minimal
- Check go.mod for current dependencies

## Project Structure

- `main.go` - CLI setup and root command
- `check.go` - Check command implementation
- `config.go` - Configuration management
- `helpers.go` - Utility functions
- `*_test.go` - Unit tests
