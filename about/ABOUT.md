# About Go Reloaded

**Version:** 1.2.1  
**Author:** Spiros Nikoloudakis  
**Email:** sp.nikoloudakis@gmail.com  
**License:** MIT License  

## What is Go Reloaded?

Go Reloaded is a text transformation tool that intelligently formats and corrects text using various linguistic and formatting rules.

## Features

- **Numeric Conversions**: Convert hexadecimal and binary numbers to decimal
- **Case Transformations**: Change text case (uppercase, lowercase, capitalize)
- **Smart Formatting**: Fix punctuation spacing and quote normalization
- **Grammar Correction**: Automatic article correction (a/an)
- **Multi-word Operations**: Apply transformations to multiple words
- **Input Validation**: Secure processing with buffer overflow protection

## How It Works

The application uses a pipeline-based architecture to process text through various transformation modules:

1. **Input Validation** - Ensures safe, keyboard-character-only input
2. **Numeric Processing** - Converts number systems
3. **Case Operations** - Applies case transformations
4. **Grammar Rules** - Corrects articles and formatting
5. **Output Generation** - Produces clean, formatted text

## System Requirements

- **Operating System**: Windows, macOS, or Linux
- **Runtime**: Go 1.21 or higher (for source code)
- **Memory**: Minimal (< 50MB)
- **Storage**: < 10MB

## Usage Modes

### Web Interface (Recommended)
- Modern browser-based interface
- Real-time validation and error highlighting
- Interactive transformation buttons
- Auto-save and session management

### Command Line Interface
- Direct file processing
- Batch operations
- Scriptable automation

## Security Features

- Buffer overflow protection (10MB limit)
- Input validation and sanitization
- Safe string operations with bounds checking
- Protection against malicious input patterns

## Support

For technical support, bug reports, or feature requests:
- **Email**: sp.nikoloudakis@gmail.com
- **Response Time**: 24-48 hours

## Copyright

© 2025 Spiros Nikoloudakis. All rights reserved.
Licensed under the MIT License.