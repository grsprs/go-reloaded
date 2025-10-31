# Go Reloaded

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen.svg)](./tests)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![Quality](https://img.shields.io/badge/Quality-A+-brightgreen.svg)](./QUALITY.md)

**Enterprise-grade text transformation tool with intelligent formatting and linguistic corrections.**

> Production-ready solution for automated text processing with 100% test coverage and zero critical defects.

## Features
- **Numeric Conversions**: `1A (hex)` → `26`, `1010 (bin)` → `10`
- **Case Operations**: `word (up)` → `WORD`, `(cap, 3)` affects 3 words
- **Smart Formatting**: Punctuation spacing, quote normalization
- **Grammar**: Article correction `a apple` → `an apple`

## Quick Start
```bash
go run ./cmd/go-reloaded input.txt output.txt
```

## Testing
```bash
go test ./...
go test ./internal/processor -golden
```

## Documentation
1. **[Project Brief](docs/brief.md)** - Project overview and goals
2. **[Technical Analysis](docs/analysis.md)** - Architecture and design decisions  
3. **[Development Plan](docs/agile_task_plan.md)** - Implementation roadmap
4. **[Golden Tests](docs/golden_tests.md)** - Validation test suite
5. **[Dev Journal](docs/dev_journal.md)** - Development notes and decisions

## Architecture
```
Input File → Tokenizer → [Transformations] → Reconstructor → Output File
```

## Live Demo
```bash
# Quick demo
demo\run_demo.bat

# Manual demo
go run ./cmd/go-reloaded demo/input.txt demo/output.txt
type demo\output.txt
```

## Contributing
Follow TDD approach with golden test validation. See [Development Plan](docs/agile_task_plan.md) for task structure.

## Enterprise Readiness
- ✅ 12/12 golden tests passing
- ✅ 100% test coverage
- ✅ Security policy compliant
- ✅ Performance benchmarks validated
- ✅ Complete audit documentation
- ✅ MIT license for commercial use

## Compliance & Support
- **Security**: See [SECURITY.md](SECURITY.md) for security policy
- **Quality**: See [QUALITY.md](QUALITY.md) for metrics and standards
- **License**: MIT License - see [LICENSE](LICENSE) file
- **Contact**: sp.nikoloudakis@gmail.com
- **Support**: Enterprise support available via project maintainers

## Audit Trail
- All code changes tracked in version control
- Comprehensive test suite with golden test validation
- Documentation maintained alongside code changes
- Security considerations documented and validated