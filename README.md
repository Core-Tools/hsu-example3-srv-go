# HSU Example3 Srv Go

HSU Repository Portability Framework - Approach 3 (Multi-Repository + Single-Language)

This is a **Go server implementation** repository for the HSU Example3 demonstration, showing how to structure a focused service implementation that depends on shared components from another repository.

## Features

- gRPC-based Echo service server implementation in Go
- Cross-platform build system with HSU Universal Makefile
- Repository-portable code structure for multi-repo scenarios
- Dependency on shared components from `hsu-example3-common`

## Related Repositories

This server implementation works with:
- `hsu-example3-common` - Shared API definitions and client implementations
- `hsu-example3-srv-py` - Python server implementation (alternative)

## Quick Start

```bash
# Build the server
make build

# Run the server
make run-srv
```

To test the server, use clients from the `hsu-example3-common` repository.

## Documentation

For comprehensive documentation, setup guides, and framework details, see:
https://github.com/Core-Tools/docs/blob/main/README.md

## Repository Structure

This repository demonstrates **Approach 3**: Multi-Repository + Single-Language, specifically the **focused service implementation** pattern that leverages shared components while maintaining a clean, single-purpose repository structure.