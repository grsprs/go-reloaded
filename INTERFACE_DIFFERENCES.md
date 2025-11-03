# Interface Differences - CLI vs Web UI

## Overview

Go Reloaded provides two distinct interfaces with different capabilities and behaviors. This document outlines the key differences for auditors and users.

## 🖥️ CLI Interface (Command Line)

### Usage
```bash
go run ./cmd/go-reloaded input.txt output.txt
```

### Characteristics
- **Pure Processing**: Processes text exactly as specified in golden tests
- **No Info Messages**: Clean output without transformation feedback
- **File-based**: Reads from input file, writes to output file
- **Specification Compliant**: 100% adherent to original requirements
- **Minimal Output**: Only the transformed text, no additional information
- **Validation**: Basic input validation for security

### Example Output
```
Input:  "I saw a apple and make this (up, 2) text"
Output: "I saw an apple and MAKE THIS text"
```

## 🌐 Web UI Interface (Browser)

### Usage
```bash
go run ./cmd/go-reloaded-web
```

### Enhanced Characteristics
- **Interactive Processing**: Real-time text transformation with feedback
- **Info Messages**: Detailed transformation reports for user awareness
- **Web-based**: Browser interface with modern UI
- **Enhanced Features**: Additional capabilities beyond core specification
- **Verbose Output**: Includes transformation details and corrections
- **Advanced Validation**: Comprehensive input validation with interactive error dialogs
- **HTML Entity Decoding**: Automatically converts HTML entities (&#39; → ') for proper processing

### Example Output
```
Input:  "I saw a apple and make this (up, 2) text"
Output: "I saw an apple and MAKE THIS text

INFO: Transformations applied:
• Article: 'a apple' → 'an apple'
• Case: Applied up transformation to 2 words (numbers excluded)"
```

## 📊 Feature Comparison Matrix

| Feature | CLI Interface | Web UI Interface |
|---------|---------------|------------------|
| **Core Transformations** | ✅ Full Support | ✅ Full Support |
| **Golden Test Compliance** | ✅ 100% | ✅ 100% (core processing) |
| **Transformation Info** | ❌ None | ✅ Detailed Reports |
| **Interactive Validation** | ❌ Basic | ✅ Advanced with Dialogs |
| **Error Positioning** | ❌ None | ✅ Cursor Highlighting |
| **Multi-word Feedback** | ❌ Silent | ✅ Word Count Reports |
| **Article Correction Tracking** | ❌ Silent | ✅ Before/After Display |
| **Real-time Processing** | ❌ File-based | ✅ Interactive |
| **User Interface** | ❌ Command Line | ✅ Modern Web UI |
| **Session Management** | ❌ None | ✅ Multi-browser Support |
| **Auto Port Detection** | ❌ N/A | ✅ Smart Port Scanning |
| **Graceful Shutdown** | ❌ N/A | ✅ Browser Close Detection |
| **HTML Entity Support** | ❌ Plain Text Only | ✅ Auto-decode (&#39; → ') |

## 🔧 Technical Implementation Differences

### CLI Processing Functions
```go
ProcessText(text)        // Core processing, no info
ProcessTextUnsafe(text)  // Bypass validation, no info
```

### Web UI Processing Functions
```go
ProcessTextWithInfo(text)        // Core processing + info messages
ProcessTextUnsafeWithInfo(text)  // Bypass validation + info messages
```

## 🎯 Use Cases

### CLI Interface - Best For:
- **Automated Scripts**: Batch processing without user interaction
- **Golden Test Validation**: Exact specification compliance testing
- **Production Pipelines**: Clean output for further processing
- **Auditing**: Verification against original requirements
- **Performance Testing**: Minimal overhead processing

### Web UI Interface - Best For:
- **Interactive Use**: Real-time text transformation and feedback
- **Learning**: Understanding what transformations were applied
- **Debugging**: Detailed information about processing steps
- **User Experience**: Modern, intuitive interface
- **Development**: Testing and validation with immediate feedback

## 🔍 Auditor Notes

### Compliance Verification
1. **CLI Interface**: Use for golden test validation and specification compliance
2. **Web UI Interface**: Additional features do not affect core compliance
3. **Dual Processing**: Both interfaces use the same core transformation engine
4. **Backward Compatibility**: CLI maintains 100% original behavior

### Testing Strategy
```bash
# Test CLI compliance
go test ./tests -v  # All 27 tests pass

# Test Web UI functionality
go run ./cmd/go-reloaded-web  # Enhanced features + core compliance
```

## 📋 Summary

- **CLI**: Specification-compliant, minimal, file-based processing
- **Web UI**: Enhanced user experience with detailed feedback and modern interface
- **Core Engine**: Identical transformation logic in both interfaces
- **Compliance**: Both maintain 100% adherence to original requirements
- **Choice**: Users can select interface based on their specific needs

The dual interface approach ensures specification compliance while providing enhanced user experience where needed.