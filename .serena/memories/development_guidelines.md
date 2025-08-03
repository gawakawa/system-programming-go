# Development Guidelines

## Code Patterns
- **Educational Focus**: Code is written for learning system programming concepts
- **Standard Library**: Prefer standard library solutions over external dependencies
- **Example-Driven**: Functions demonstrate specific I/O patterns and techniques

## Naming Conventions
- Mix of camelCase and snake_case (historical reasons)
- Descriptive function names indicating their purpose
- Chapter-based organization for related examples

## I/O Patterns
- **Chapter 02**: Focus on io.Writer interface implementations
- **Chapter 03**: Focus on io.Reader interface implementations
- Demonstrate real-world use cases (web, files, compression, etc.)

## Design Principles
- Keep examples self-contained and runnable
- Demonstrate one concept per function
- Use practical scenarios (PNG manipulation, web servers, file operations)
- Minimal external dependencies

## Development Flow
1. Write code following Go standards
2. Test execution in appropriate chapter directory
3. Format with `go fmt`
4. Verify with `go vet`
5. Commit changes with descriptive messages