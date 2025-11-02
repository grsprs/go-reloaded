# Changelog

All notable changes to this project will be documented in this file.

## [1.2.2] - 2025-11-01

### Added
- **Comprehensive Bracket Validation**: All keyboard brackets (parentheses, square, curly, angle)
- **Enhanced Quote Detection**: Double quotes, single quotes with contraction handling
- **Sequential Error Dialogs**: Individual dialogs for each syntax error with cursor positioning
- **Intentional Error Override**: Checkbox system to bypass validation for intentional syntax

### Enhanced
- **Error Detection**: Covers all opening/closing keyboard characters
- **User Experience**: Cursor jumps to exact error position for each issue
- **Validation Logic**: Smart contraction detection (skips "don't", "can't", etc.)
- **Error Messages**: Clear, actionable guidance for each bracket/quote type

### Technical
- **Client-side Validation**: Real-time error detection before server submission
- **Safe Processing**: ProcessTextUnsafe() function for intentional errors
- **Multi-error Handling**: Sequential processing of all syntax issues

## [1.2.1] - 2025-11-01

### Added
- **Keyboard Character Validation**: Ensures only standard keyboard characters are accepted
- **Binary Content Detection**: Prevents processing of non-text content (images, files, etc.)
- **Exit Button**: Clean application exit with confirmation dialog
- **Enhanced Error Display**: Better error messages with context highlighting
- **Client-side Validation**: Real-time detection of non-keyboard and binary content

### Enhanced
- **Character Filtering**: Blocks emojis, Unicode symbols, and non-keyboard characters
- **User Experience**: Clear warnings with character position when invalid content detected
- **Error Handling**: Improved error display in web interface
- **Input Guidance**: Updated help text to clarify keyboard character requirements

## [1.2.0] - 2025-11-01

### Added
- **Input Validation**: Buffer overflow protection (10MB limit)
- **Syntax Checking**: Detects unclosed brackets and quotes with error highlighting
- **Safe Operations**: Bounds-checked string operations to prevent crashes
- **Security**: Protection against malformed input and edge cases
- **Error Messages**: Clear validation errors with context highlighting

### Enhanced
- **Quote Handling**: Smart detection of contractions vs. actual quotes
- **Memory Safety**: All string operations use safe bounds checking
- **Robustness**: Graceful handling of invalid input patterns

## [1.1.3] - 2025-11-01

### Fixed
- **Documentation Consistency**: Unified version numbers across all files (v1.1.3)
- **Date Consistency**: Updated all dates to November 1, 2025
- **Content Cleanup**: Removed redundant "comprehensive" and "professional" language
- **Version Display**: Updated web interface footer to show correct version

## [1.1.2] - 2025-11-01

### Added
- **Multi-Browser Session Tracking**: Server now tracks multiple browser sessions
- **Smart Shutdown Logic**: Server shuts down when ANY browser window is closed
- **Session Management**: Proper cleanup of browser sessions on window close

## [1.1.1] - 2025-11-01

### Fixed
- **UI Color Scheme**: Updated background color to dark grey (#4a5568) for better visual contrast
- **Button Styling**: Darkened button gradient colors (#4a4fdf, #8b4cdf) for improved consistency
- **Visual Polish**: Minor UI refinements for improved appearance

## [1.1.0] - 2025-11-01

### Added
- **Web Interface**: Modern web interface with clean design
- **Auto Port Detection**: Automatically finds available ports (8080+) to prevent conflicts
- **Smart Shutdown**: Server automatically shuts down when browser window closes
- **Dark Mode Support**: Automatic detection and support for system dark mode
- **Interactive Transformation Buttons**: Click-to-insert examples for all transformations
- **Real-time Processing**: Live text transformation with form submission
- **Responsive Design**: Optimized for 1080p displays without scrolling
- **Graceful Error Handling**: Proper port conflict resolution and resource cleanup

### Enhanced
- **CLI Interface**: Added usage instructions for both CLI and Web modes
- **Documentation**: Updated README with web UI information
- **Build System**: Added Makefile for standardized build commands
- **CI/CD Pipeline**: GitHub Actions for automated testing and releases

### Technical Improvements
- **Cross-platform Browser Opening**: Automatic browser launch on Windows/macOS/Linux
- **Resource Management**: Proper cleanup of network resources and goroutines
- **Port Management**: Background port availability checking
- **Template Security**: Safe HTML template rendering with proper escaping

## [1.0.0] - 2025-10-31

### Added
- **Core Text Processing**: Complete implementation of all transformation requirements
- **Numeric Conversions**: Hexadecimal and binary to decimal conversion
- **Case Transformations**: Uppercase, lowercase, and capitalize operations
- **Multi-word Support**: Transform multiple words with count syntax
- **Punctuation Formatting**: Smart spacing and grouping
- **Quote Normalization**: Proper quote spacing and formatting
- **Article Correction**: Intelligent a/an correction using phonetic rules
- **Golden Test Suite**: 12 comprehensive tests covering all requirements
- **Complete Documentation**: Documentation suite
- **Security Policy**: Comprehensive security guidelines and practices
- **Quality Metrics**: Code quality standards and compliance documentation

### Technical Features
- **Pipeline Architecture**: Modular transformation pipeline for maintainability
- **Pattern-based Rules**: Generic rules instead of hardcoded word lists
- **Error Handling**: Comprehensive error handling and validation
- **Performance Optimization**: Sub-second execution time for all operations
- **Cross-platform Support**: Windows, macOS, and Linux compatibility

### Documentation
- **Project Brief**: Technical overview and architecture decisions
- **Development Plan**: Agile task breakdown and implementation roadmap
- **Golden Tests Specification**: Comprehensive test validation approach
- **Development Journal**: Detailed development notes and decisions
- **Audit Navigation**: Auditor guidance and checklists