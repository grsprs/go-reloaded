# Go Reloaded - Help Guide

## Quick Start

### Web Interface
1. Open the application (it will launch automatically in your browser)
2. Type or paste text in the input area
3. Click transformation buttons or type commands directly
4. Press "Transform Text" or double-press Enter to process
5. View results in the output area

### Command Line
```bash
go run ./cmd/go-reloaded input.txt output.txt
```

## Transformation Commands

### Number Conversions
- `42 (hex)` → Converts hexadecimal 42 to decimal 66
- `1010 (bin)` → Converts binary 1010 to decimal 10

### Case Operations
- `word (up)` → Converts to UPPERCASE
- `WORD (low)` → Converts to lowercase  
- `word (cap)` → Capitalizes First Letter
- `(up, 3)` → Applies uppercase to next 3 words

### Formatting
- Automatic punctuation spacing: `word ,` → `word,`
- Quote normalization: `' text '` → `'text'`
- Article correction: `a apple` → `an apple`

## Input Guidelines

### Supported Characters
- Standard keyboard characters (A-Z, 0-9, symbols)
- Basic punctuation and whitespace
- Extended ASCII for international keyboards

### Not Supported
- Emojis (😀, 🎉)
- Unicode symbols (∑, ∞)
- Non-keyboard characters
- Binary or file content

## Error Messages

### Common Errors
- **"Non-keyboard character detected"**: Remove emojis or special Unicode
- **"Unclosed quote/bracket"**: Check for matching quotes and parentheses
- **"Buffer overflow"**: Input too large (max 10MB)
- **"Binary content detected"**: Paste only plain text

## Keyboard Shortcuts

### Web Interface
- **Enter**: New line
- **Double Enter**: Transform text (within 500ms)
- **Ctrl+A**: Select all text
- **Ctrl+C/V**: Copy/paste

## Tips & Tricks

1. **Batch Processing**: Use multiple commands in one input
2. **Preview Mode**: See transformations before applying
3. **Error Recovery**: Fix highlighted issues before processing
4. **Session Persistence**: Your work is saved during the session

## Troubleshooting

### Application Won't Start
- Check if port 8080+ is available
- Ensure Go runtime is installed (for source)
- Try running as administrator

### Transformations Not Working
- Check input for unsupported characters
- Verify command syntax: `text (command)`
- Ensure quotes and brackets are properly closed

### Performance Issues
- Reduce input size (max 10MB recommended)
- Close other browser tabs
- Restart the application

## Getting Help

If you need additional assistance:
1. Check error messages for specific guidance
2. Review input for unsupported content
3. Contact support: sp.nikoloudakis@gmail.com

## Version Information

Current version includes all transformation features with enhanced security and validation.