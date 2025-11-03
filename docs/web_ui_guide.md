# Web UI User Guide

## Overview

Go Reloaded v1.4.0 includes a modern web interface with comprehensive input validation, interactive error handling, and earth-tone design that provides an intuitive experience for text transformation operations.

## Quick Start

```bash
go run ./cmd/go-reloaded-web
```

The web interface will:
1. **Auto-detect** an available port (starting from 8080)
2. **Launch your browser** automatically
3. **Display** the interface
4. **Auto-shutdown** when you close the browser

## Interface Layout

### Header
- **Logo**: "GR" branded icon with "Go Reloaded" title
- **About Menu**: Access to documentation, help, changelog, and license
- **Exit Button**: Graceful shutdown with confirmation dialog

### Three-Column Layout

#### Left Panel - Input Transformations
- **Numeric Conversions**
  - Hexadecimal → Decimal
  - Binary → Decimal
- **Case Transformations**
  - Uppercase
  - Lowercase
  - Capitalize

#### Center Panel - Text Processing
- **Input Textarea**: Enter or paste your text
- **Output Textarea**: View transformed results
- **Transform Button**: Process the text

#### Right Panel - Formatting & Grammar
- **Text Formatting**
  - Fix Punctuation
  - Normalize Quotes
- **Grammar Correction**
  - Article Correction (a/an)
  - Multi-word Transform

## Features

### Interactive Buttons
- **Click any button** to insert example text into the input area
- **Examples shown** below each button demonstrate the transformation
- **Cursor positioning** maintained for precise text insertion

### Real-time Processing
- **Form submission** processes all transformations
- **Input validation** with interactive error dialogs
- **Sequential error handling** with cursor positioning
- **Auto-clear** output when input is cleared
- **Instant feedback** on transformation results
- **Double Enter** shortcut for quick transformation

### Modern Design
- **Earth-tone color scheme** with warm, professional styling
- **Dark mode support** automatically detects system preference
- **Responsive layout** optimized for 1080p displays
- **No scrolling required** - fits perfectly in viewport
- **Interactive error dialogs** with cursor highlighting
- **Smooth animations** and hover effects

### Smart Port Management
- **Auto-detection** finds available ports (8080, 8081, 8082, etc.)
- **Conflict resolution** prevents "port already in use" errors
- **Clean startup** every time

### Graceful Shutdown
- **Browser close detection** automatically shuts down server
- **Resource cleanup** frees ports and memory
- **Manual shutdown** available with Ctrl+C

## Usage Examples

### Basic Text Transformation
1. Type or paste text in the input area
2. Click "Transform Text" button
3. View results in the output area

### Using Quick Insert Buttons
1. Click any transformation button (e.g., "Hex to Decimal")
2. Example text is inserted at cursor position
3. Click "Transform Text" to see results

### Multiple Transformations
1. Insert multiple transformation examples
2. Single click processes all transformations in sequence
3. Output shows cumulative results

## Technical Details

### Browser Compatibility
- **Modern browsers** with HTML5 and CSS3 support
- **JavaScript required** for interactive features
- **Responsive design** works on desktop and tablet

### Performance
- **Instant startup** typically under 1 second
- **Fast processing** all transformations complete in milliseconds
- **Low resource usage** minimal CPU and memory footprint

### Security
- **Local processing** all text processing happens locally
- **No data transmission** text never leaves your machine
- **Input validation** comprehensive security checks
- **Buffer overflow protection** safe string operations
- **Bracket/quote validation** with interactive error handling
- **Safe shutdown** proper cleanup of all resources
- **HTML entity decoding** automatically converts &#39; to ' for proper quote processing
- **Enhanced exit button** with improved reliability and user feedback
- **Always-on formatting** automatic punctuation and quote processing with feedback
- **Enhanced bracket support** comprehensive handling of all bracket types
- **Real-time feedback** green info area shows applied transformations

## Troubleshooting

### Port Issues
- **Auto-detection** handles port conflicts automatically
- **Range scanning** tries ports 8080-8179
- **Fallback behavior** uses original port if all busy

### Browser Issues
- **Manual access** navigate to displayed URL if auto-open fails
- **Refresh page** if interface doesn't load properly
- **Check console** for JavaScript errors in browser dev tools

### Server Issues
- **Restart server** stop with Ctrl+C and restart
- **Check permissions** ensure Go has network access
- **Firewall settings** may block local server access

## Input Validation Features

### Interactive Error Handling
- **Sequential validation** checks all bracket and quote pairs
- **Cursor positioning** highlights exact error location
- **Error dialogs** provide clear descriptions and options
- **Intentional override** allows processing with syntax errors
- **Buffer overflow protection** prevents malicious input

### Supported Validations
- **Bracket matching**: (), [], {}, <>
- **Quote pairing**: single and double quotes
- **Keyboard characters only** prevents binary content
- **Buffer size limits** protects against overflow attacks
- **Smart contraction detection** handles "don't", "can't", etc.

### Error Dialog Options
- **Fix It**: Return to input for manual correction
- **Continue**: Process text with intentional syntax errors
- **Checkbox**: Mark errors as intentional for batch processing

## Advanced Usage

### Command Line Options
```bash
# Standard startup
go run ./cmd/go-reloaded-web

# Build and run
make build-web
./bin/go-reloaded-web

# Quick run target
make run-web
```

### Integration
- **API endpoint** available at root path for programmatic access
- **Form submission** standard HTTP POST to root
- **JSON response** available for automated processing

## Support

For technical support or feature requests:
- **Email**: sp.nikoloudakis@gmail.com
- **Documentation**: See project README and technical docs
- **Issues**: Report via project repository