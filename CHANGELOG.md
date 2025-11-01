# Changelog

All notable changes to this project will be documented in this file.

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
- **Documentation**: Updated README with comprehensive web UI information
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
- **Complete Documentation**: Comprehensive documentation suite
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