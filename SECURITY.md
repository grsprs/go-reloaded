# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.0.x   | :white_check_mark: |

## Security Considerations

### Input Validation
- All file inputs are validated for size and format
- Regex patterns are bounded to prevent ReDoS attacks
- No external dependencies that could introduce vulnerabilities

### Data Processing
- No sensitive data is logged or persisted
- All transformations are performed in memory
- No network connections or external API calls

### File Operations
- File permissions are restricted (0644 for output files)
- No arbitrary file path execution
- Input/output paths are validated

## Reporting a Vulnerability

If you discover a security vulnerability, please report it via:
- Email: sp.nikoloudakis@gmail.com
- Create a private issue in the repository

**Response Time:** 48 hours for acknowledgment, 7 days for resolution.