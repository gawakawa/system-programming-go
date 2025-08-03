# Task Completion Checklist

## Code Quality Checks
1. **Format code**: `go fmt ./...`
2. **Static analysis**: `go vet ./...`
3. **Build check**: `go build ./...`

## Testing (when applicable)
- Currently no tests exist in the project
- If tests are added, run: `go test ./...`

## Execution Verification
- Test running the code in appropriate chapter directory
- For chapter02: `cd chapter02 && go run main.go` (web server)
- For chapter03: `cd chapter03 && go run main.go` (may need input files)

## Git Workflow
- Check status: `git status`
- Stage changes: `git add .`
- Commit with descriptive message: `git commit -m "description"`

## Environment Considerations
- Ensure Nix environment is active if using flake
- Check that Go version is compatible (1.24+)
- Verify any required input files exist for examples