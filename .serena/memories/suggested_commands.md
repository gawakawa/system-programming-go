# Suggested Commands

## Development Commands
- `go run main.go` - Run code in specific chapter directory
- `go build` - Build the current module
- `go vet ./...` - Static analysis tool to find potential issues
- `go fmt ./...` - Format all Go source files
- `go test ./...` - Run tests (currently no test files exist)

## Environment Setup
- `nix develop` - Enter development shell with Go and gopls
- `direnv allow` - Allow direnv to load environment variables

## Chapter-Specific Execution
- `cd chapter02 && go run main.go` - Runs web server on :8080
- `cd chapter03 && go run main.go` - Runs file I/O examples (may need input files)

## System Commands (Darwin/macOS)
- `ls -la` - List files with details
- `find . -name "*.go"` - Find Go files
- `grep -r "pattern" .` - Search for patterns in files
- `git status` - Check git status
- `git add .` - Stage all changes
- `git commit -m "message"` - Commit changes