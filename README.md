# Go Reloaded

Text transformation tool with intelligent formatting and linguistic corrections.

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

## Contributing
Follow TDD approach with golden test validation. See [Development Plan](docs/agile_task_plan.md) for task structure.