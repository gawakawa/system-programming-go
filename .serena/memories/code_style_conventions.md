# Code Style and Conventions

## Go Standards
- Follows standard Go formatting (gofmt)
- Uses standard Go naming conventions:
  - Functions: camelCase (e.g., `copyFile`, `writeBuffer`)
  - No type annotations (Go uses type inference)
  - Standard Go project structure

## Code Organization
- Each chapter has its own directory (chapter02/, chapter03/)
- Single main.go file per chapter containing multiple example functions
- Functions are typically standalone examples demonstrating specific concepts
- Main function calls specific example functions

## Naming Patterns
- Functions often prefixed by their purpose:
  - `write_*` for writer examples (e.g., `write_file`, `write_console`)
  - `read*` for reader examples (e.g., `readChunks`, `ReadFile`)
- Underscore naming used for some functions (non-standard but consistent within project)

## No Documentation Requirements
- No docstrings or comments observed in the codebase
- Focus on educational/example code rather than production code