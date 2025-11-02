// Copyright (c) 2025 Spiros Nikoloudakis
// Licensed under MIT License - see LICENSE file for details

package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"go-reloaded/internal/processor"
)

type PageData struct {
	Input  string
	Output string
	Error  string
}

var (
	activeSessions = make(map[string]bool)
	sessionMutex   sync.Mutex
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Go Reloaded - Text Transformation Suite</title>
<style>
:root {
  --color-bg: #8b7355;
  --color-bg-panel: #f5f1eb;
  --color-text: #3c2e26;
  --color-accent: #a0845c;
  --color-accent-gradient: linear-gradient(135deg, #a0845c, #8b7355);
  --color-border: #d4c4a8;
  --color-shadow: rgba(60, 46, 38, 0.1);
  --radius: 10px;
  --transition: 0.2s ease;
  --font-sans: "Inter", "Segoe UI", sans-serif;
  --font-mono: "Fira Code", monospace;
}

@media (prefers-color-scheme: dark) {
  :root {
    --color-bg: #2c1810;
    --color-bg-panel: #3d2817;
    --color-text: #e8ddd4;
    --color-accent: #d4a574;
    --color-accent-gradient: linear-gradient(135deg, #d4a574, #b8935f);
    --color-border: #4a3426;
    --color-shadow: rgba(0, 0, 0, 0.6);
  }
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: var(--font-sans);
  background: var(--color-bg);
  color: var(--color-text);
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: hidden;
}

header {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  background: var(--color-bg-panel);
  border-bottom: 1px solid var(--color-border);
  box-shadow: 0 1px 2px var(--color-shadow);
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.dropdown {
  position: relative;
  display: inline-block;
}

.dropdown-btn {
  background: var(--color-accent-gradient);
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px 12px;
  font-size: 0.9rem;
  cursor: pointer;
  transition: var(--transition);
}

.dropdown-btn:hover {
  opacity: 0.9;
}

.dropdown-content {
  display: none;
  position: absolute;
  right: 0;
  background: var(--color-bg-panel);
  min-width: 120px;
  box-shadow: 0 4px 8px var(--color-shadow);
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  z-index: 1000;
}

.dropdown-content a {
  color: var(--color-text);
  padding: 8px 12px;
  text-decoration: none;
  display: block;
  font-size: 0.9rem;
  transition: var(--transition);
}

.dropdown-content a:hover {
  background: var(--color-accent);
  color: #fff;
}

.dropdown:hover .dropdown-content {
  display: block;
}

.exit-btn-header {
  background: #b85450;
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px 16px;
  font-weight: 500;
  font-size: 0.9rem;
  cursor: pointer;
  transition: var(--transition);
}

.exit-btn-header:hover {
  background: #a04844;
  transform: translateY(-1px);
}

header .logo {
  display: flex;
  align-items: center;
  font-weight: 600;
  font-size: 1.2rem;
}

header .logo span.icon {
  background: var(--color-accent-gradient);
  color: #fff;
  font-weight: bold;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 10px;
}

header .subtitle {
  font-size: 0.95rem;
  opacity: 0.8;
}

main {
  flex: 1;
  display: grid;
  grid-template-columns: 280px 1fr 280px;
  grid-template-rows: 1fr;
  gap: 20px;
  padding: 20px;
  background: var(--color-bg);
}

aside {
  background: var(--color-bg-panel);
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  box-shadow: 0 2px 4px var(--color-shadow);
  padding: 16px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  gap: 24px;
}

aside h3 {
  font-size: 0.95rem;
  margin-bottom: 8px;
  opacity: 0.8;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

aside .button {
  background: var(--color-accent-gradient);
  color: #fff;
  border: none;
  padding: 10px 12px;
  border-radius: var(--radius);
  font-size: 0.9rem;
  text-align: left;
  cursor: pointer;
  transition: var(--transition);
  box-shadow: 0 2px 4px var(--color-shadow);
}

aside .button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

aside .example {
  font-size: 0.8rem;
  opacity: 0.7;
  margin-top: 4px;
  margin-bottom: 16px;
  margin-left: 4px;
  font-family: var(--font-mono);
}

section.center {
  display: flex;
  flex-direction: column;
  background: var(--color-bg-panel);
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  box-shadow: 0 2px 4px var(--color-shadow);
  padding: 16px;
}

textarea {
  width: 100%;
  height: 120px;
  resize: none;
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  padding: 12px;
  font-family: var(--font-mono);
  font-size: 0.95rem;
  line-height: 1.4;
  background: var(--color-bg);
  color: var(--color-text);
  transition: var(--transition);
}

textarea:focus {
  border-color: var(--color-accent);
  outline: none;
  box-shadow: 0 0 0 2px rgba(90, 95, 255, 0.2);
}

.output {
  margin-top: 12px;
  background: var(--color-bg);
}

.button-row {
  margin-top: 16px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
}

.transform-btn {
  background: var(--color-accent-gradient);
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 12px 32px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: var(--transition);
}

.transform-btn:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.clear-btn {
  background: #b85450;
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px 20px;
  font-weight: 500;
  font-size: 0.9rem;
  cursor: pointer;
  transition: var(--transition);
}

.clear-btn:hover {
  background: #a04844;
  transform: translateY(-1px);
}

.error-message {
  margin-top: 12px;
  padding: 12px;
  background: #f4e6d7;
  border: 1px solid #e6d3c1;
  border-radius: var(--radius);
  color: #8b4513;
  font-size: 0.9rem;
  white-space: pre-wrap;
  display: none;
}

@media (prefers-color-scheme: dark) {
  .error-message {
    background: #4a2c1a;
    border: 1px solid #5c3a26;
    color: #e8b894;
  }
}

.how-to {
  margin-top: 12px;
  padding: 12px;
  background: rgba(160, 132, 92, 0.15);
  border-radius: var(--radius);
  font-size: 0.85rem;
  line-height: 1.4;
}

@media (prefers-color-scheme: dark) {
  .how-to {
    background: rgba(212, 165, 116, 0.1);
  }
}

.error-dialog {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  z-index: 2000;
  display: flex;
  align-items: center;
  justify-content: center;
}

.error-dialog-content {
  background: var(--color-bg-panel);
  border: 1px solid var(--color-border);
  border-radius: var(--radius);
  padding: 20px;
  max-width: 400px;
  width: 90%;
  box-shadow: 0 4px 12px var(--color-shadow);
}

.error-dialog h4 {
  margin-bottom: 12px;
  color: var(--color-accent);
  font-size: 1.1rem;
}

.checkbox-container {
  margin: 16px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.checkbox-container input[type="checkbox"] {
  width: 16px;
  height: 16px;
}

.dialog-buttons {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  margin-top: 16px;
}

.proceed-btn {
  background: var(--color-accent-gradient);
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px 16px;
  cursor: pointer;
  transition: var(--transition);
}

.cancel-btn {
  background: #b85450;
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px 16px;
  cursor: pointer;
  transition: var(--transition);
}

.proceed-btn:hover, .cancel-btn:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

.cursor-highlight {
  background-color: rgba(255, 0, 0, 0.3);
  animation: blink 1s infinite;
}

@keyframes blink {
  0%, 50% { background-color: rgba(255, 0, 0, 0.3); }
  51%, 100% { background-color: transparent; }
}

.how-to h4 {
  margin-bottom: 8px;
  color: var(--color-accent);
  font-size: 0.9rem;
}

footer {
  height: 40px;
  background: var(--color-bg-panel);
  border-top: 1px solid var(--color-border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  font-size: 0.85rem;
  opacity: 0.8;
}

@media (max-width: 1500px) {
  main {
    grid-template-columns: 240px 1fr 240px;
    gap: 16px;
  }
}
</style>
</head>
<body>

<header>
  <div class="logo">
    <span class="icon">GR</span> Go Reloaded
  </div>
  <div class="header-right">
    <div class="dropdown">
      <button class="dropdown-btn">About ▼</button>
      <div class="dropdown-content">
        <a href="#" onclick="showAbout()">About</a>
        <a href="#" onclick="showHelp()">Help</a>
        <a href="#" onclick="showChangelog()">What's New</a>
        <a href="#" onclick="showLicense()">License</a>
      </div>
    </div>
    <button class="exit-btn-header" onclick="exitApp()">Exit</button>
  </div>
</header>

<main>
  <aside>
    <div>
      <h3>Numbers</h3>
      <button class="button" onclick="insertText('(hex)')">Hex → Decimal</button>
      <div class="example">2A (hex) → 42</div>
      <button class="button" onclick="insertText('(bin)')">Binary → Decimal</button>
      <div class="example">1010 (bin) → 10</div>
    </div>
    <div>
      <h3>Case</h3>
      <button class="button" onclick="insertText('(up)')">Uppercase</button>
      <div class="example">hello (up) → HELLO</div>
      <button class="button" onclick="insertText('(low)')">Lowercase</button>
      <div class="example">HELLO (low) → hello</div>
      <button class="button" onclick="insertText('(cap)')">Capitalize</button>
      <div class="example">hello (cap) → Hello</div>
    </div>
  </aside>

  <section class="center">
    <form method="POST">
      <textarea name="input" id="input" placeholder="Enter your text here...">{{.Input}}</textarea>
      <textarea id="output" class="output" placeholder="Transformed output..." readonly>{{.Output}}</textarea>
      <div class="button-row">
        <button type="submit" class="transform-btn">Transform Text</button>
        <button type="button" class="clear-btn" onclick="return clearText(event);">Clear</button>
      </div>
    </form>
    <div class="error-message" {{if .Error}}style="display: block;"{{end}}>{{.Error}}</div>
    <div class="error-dialog" id="errorDialog" style="display: none;">
      <div class="error-dialog-content">
        <h4>Issue Found</h4>
        <p id="errorText"></p>
        <div class="checkbox-container">
          <input type="checkbox" id="intentionalCheck">
          <label for="intentionalCheck">This is intentional - continue processing</label>
        </div>
        <div class="dialog-buttons">
          <button onclick="proceedWithError()" class="proceed-btn">Continue</button>
          <button onclick="closeErrorDialog()" class="cancel-btn">Fix It</button>
        </div>
      </div>
    </div>
    <div class="how-to">
      <h4>How to Use</h4>
      • Click buttons to insert transformations<br>
      • Single Enter: New line<br>
      • Double Enter: Transform text<br>
      • Use Transform button or double Enter<br>
      • Only plain text is supported
    </div>
  </section>

  <aside>
    <div>
      <h3>Format</h3>
      <button class="button" onclick="insertText(' ,')">Fix Punctuation</button>
      <div class="example">hello , world → hello, world</div>
      <button class="button" onclick="insertText(&quot;' text '&quot;)">Normalize Quotes</button>
      <div class="example">' hello world ' → 'hello world'</div>
    </div>
    <div>
      <h3>Grammar</h3>
      <button class="button" onclick="insertText('a apple')">Article Correction</button>
      <div class="example">I saw a apple → I saw an apple</div>
      <button class="button" onclick="insertText('(up, 2)')">Multi-word Transform</button>
      <div class="example">hello world (up, 2) → HELLO WORLD</div>
    </div>
  </aside>
</main>

<footer>
  <span>Status: Ready</span>
  <span>Text Processing Engine v1.2.2</span>
</footer>

<script>
  function insertText(text) {
    const input = document.getElementById('input');
    const start = input.selectionStart;
    const end = input.selectionEnd;
    const value = input.value;
    
    input.value = value.substring(0, start) + text + ' ' + value.substring(end);
    input.focus();
    input.setSelectionRange(start + text.length + 1, start + text.length + 1);
  }
  
  function clearText(event) {
    if (event) {
      event.preventDefault();
      event.stopPropagation();
    }
    document.getElementById('input').value = '';
    document.getElementById('output').value = '';
    document.getElementById('input').focus();
    hideError();
    return false;
  }
  
  function exitApp() {
    if (confirm('Are you sure you want to exit the application?')) {
      fetch('/shutdown', { method: 'POST' }).catch(() => {});
      window.close();
    }
  }
  
  function showError(message) {
    const errorDiv = document.querySelector('.error-message');
    if (errorDiv) {
      errorDiv.textContent = message;
      errorDiv.style.display = 'block';
    }
  }
  
  function hideError() {
    const errorDiv = document.querySelector('.error-message');
    if (errorDiv) {
      errorDiv.style.display = 'none';
    }
  }
  
  let lastEnterTime = 0;
  
  document.getElementById('input').addEventListener('keydown', function(e) {
    if (e.key === 'Enter') {
      const currentTime = Date.now();
      if (currentTime - lastEnterTime < 500) {
        // Double press - transform
        e.preventDefault();
        document.querySelector('form').submit();
        lastEnterTime = 0;
      } else {
        // Single press - allow newline
        lastEnterTime = currentTime;
      }
    }
  });
  
  document.getElementById('input').addEventListener('input', function() {
    if (this.value.trim() === '') {
      document.getElementById('output').value = '';
      hideError();
    }
    
    // Check for obvious binary content and non-keyboard characters
    const text = this.value;
    if (text.includes('\x00') || /[\x00-\x08\x0E-\x1F\x7F-\x9F]/.test(text)) {
      showError('Warning: Binary or non-text content detected. Please paste only plain text.');
    }
    
    // Check for non-keyboard characters
    for (let i = 0; i < text.length; i++) {
      const char = text[i];
      const code = char.charCodeAt(0);
      
      // Allow standard keyboard characters
      if ((code >= 32 && code <= 126) || // Printable ASCII
          (code >= 160 && code <= 255) || // Extended ASCII
          char === '\t' || char === '\n' || char === '\r' || char === ' ') {
        continue;
      }
      
      // Allow common Unicode symbols
      const commonSymbols = ['€', '£', '¥', '©', '®', '™', '°'];
      if (commonSymbols.includes(char)) {
        continue;
      }
      
      showError('Warning: Non-keyboard character detected: \'' + char + '\' at position ' + (i+1) + '. Please use only standard keyboard characters.');
      break;
    }
  });
  
  // Handle form submission with error checking
  document.querySelector('form').addEventListener('submit', function(e) {
    if (intentionalError) {
      intentionalError = false;
      return true; // Allow submission
    }
    
    const input = document.getElementById('input');
    const text = input.value;
    
    // Check for validation errors before submission
    const allErrors = getAllErrors(text);
    if (allErrors.length > 0) {
      e.preventDefault();
      showAllErrors(allErrors);
      return false;
    }
    
    hideError();
  });
  

  
  function getAllErrors(text) {
    const errors = [];
    
    // Define all bracket pairs
    const bracketPairs = [
      { open: '(', close: ')', name: 'parenthesis', openName: 'opening parenthesis', closeName: 'closing parenthesis' },
      { open: '[', close: ']', name: 'square bracket', openName: 'opening square bracket', closeName: 'closing square bracket' },
      { open: '{', close: '}', name: 'curly brace', openName: 'opening curly brace', closeName: 'closing curly brace' },
      { open: '<', close: '>', name: 'angle bracket', openName: 'opening angle bracket', closeName: 'closing angle bracket' }
    ];
    
    // Check each bracket type
    for (let bracket of bracketPairs) {
      let stack = [];
      for (let i = 0; i < text.length; i++) {
        if (text[i] === bracket.open) {
          stack.push(i);
        } else if (text[i] === bracket.close) {
          if (stack.length === 0) {
            errors.push({
              message: 'You have a ' + bracket.closeName + ' (' + bracket.close + ') without a matching ' + bracket.openName + '. Please add the ' + bracket.openName + ' or remove the extra ' + bracket.closeName + '.',
              position: i,
              type: 'UNMATCHED_BRACKET'
            });
          } else {
            stack.pop();
          }
        }
      }
      
      // Check for unclosed opening brackets
      for (let pos of stack) {
        errors.push({
          message: 'You have an unclosed ' + bracket.openName + ' (' + bracket.open + ') in your text. Please add the ' + bracket.closeName + ' or remove it if not needed.',
          position: pos,
          type: 'UNCLOSED_BRACKET'
        });
      }
    }
    
    // Check for unclosed quotes and special characters
    const quoteTypes = [
      { char: '"', name: 'double quote' },
      { char: "\'", name: 'single quote', skipContractions: true }
    ];
    
    for (let quoteType of quoteTypes) {
      let positions = [];
      
      for (let i = 0; i < text.length; i++) {
        if (text[i] === quoteType.char) {
          // Skip contractions for single quotes
          if (quoteType.skipContractions && i > 0 && i < text.length - 1 && 
              /[a-zA-Z]/.test(text[i-1]) && /[a-zA-Z]/.test(text[i+1])) {
            continue;
          }
          
          if (positions.length > 0 && positions[positions.length - 1] !== -1) {
            positions[positions.length - 1] = -1; // Mark as closed
          } else {
            positions.push(i);
          }
        }
      }
      
      // Add unclosed quote errors
      for (let pos of positions) {
        if (pos !== -1) {
          errors.push({
            message: 'You have an unclosed ' + quoteType.name + ' (' + quoteType.char + ') in your text. Please add the closing ' + quoteType.name + ' or remove it if not needed.',
            position: pos,
            type: 'UNCLOSED_QUOTE'
          });
        }
      }
    }
    
    // Sort errors by position
    errors.sort((a, b) => a.position - b.position);
    
    return errors;
  }
  
  // Track if we're intentionally staying on the page
  let stayingOnPage = false;
  
  // Mark that we're staying when clicking buttons
  document.addEventListener('click', function(e) {
    if (e.target.type === 'button' || e.target.tagName === 'BUTTON') {
      stayingOnPage = true;
      setTimeout(() => { stayingOnPage = false; }, 100);
    }
  });
  
  // Only send shutdown on actual page unload
  window.addEventListener('beforeunload', function(e) {
    if (!stayingOnPage) {
      navigator.sendBeacon('/shutdown');
    }
  });
  
  function showAbout() {
    alert('Go Reloaded v1.2.1\n\nText transformation tool with intelligent formatting.\n\nAuthor: Spiros Nikoloudakis\nEmail: sp.nikoloudakis@gmail.com\nLicense: MIT License\n\nFeatures:\n• Numeric conversions (hex/bin to decimal)\n• Case transformations (up/low/cap)\n• Smart formatting and grammar correction\n• Input validation and security protection');
  }
  
  function showHelp() {
    alert('Go Reloaded - Quick Help\n\nTransformation Commands:\n• 42 (hex) → Convert hex to decimal\n• 1010 (bin) → Convert binary to decimal\n• word (up) → UPPERCASE\n• WORD (low) → lowercase\n• word (cap) → Capitalize\n• (up, 3) → Apply to 3 words\n\nFormatting:\n• Automatic punctuation spacing\n• Quote normalization\n• Article correction (a/an)\n\nInput: Only keyboard characters allowed\nShortcuts: Double Enter = Transform');
  }
  
  function showChangelog() {
    alert('What\'s New in v1.2.1\n\n✨ New Features:\n• Enhanced security - keyboard characters only\n• Better error messages with position highlighting\n• Real-time input validation\n• Earth tone theme with improved dark mode\n\n🔒 Security:\n• Buffer overflow protection (10MB limit)\n• Input validation for quotes/brackets\n• Safe string operations\n\n🎨 Interface:\n• Warm earth tone colors\n• Better button layout\n• Improved accessibility');
  }
  
  function showLicense() {
    alert('Go Reloaded License\n\nMIT License © 2025 Spiros Nikoloudakis\n\n✅ You CAN:\n• Use for any purpose (personal/commercial)\n• Copy and distribute\n• Modify the source code\n• Include in your projects\n\n❌ You CANNOT:\n• Remove copyright notices\n• Hold us liable for damages\n\nThis software is FREE to use for everyone.\nNo fees, no registration required.');
  }
  
  let errorPosition = -1;
  let intentionalError = false;
  let currentErrors = [];
  let currentErrorIndex = 0;
  
  function showErrorDialog(message, position) {
    errorPosition = position;
    document.getElementById('errorText').textContent = message;
    document.getElementById('intentionalCheck').checked = false;
    document.getElementById('errorDialog').style.display = 'flex';
    
    // Position cursor at error location
    if (position >= 0) {
      const input = document.getElementById('input');
      input.focus();
      input.setSelectionRange(position, position + 1);
      
      // Highlight the error position
      setTimeout(() => {
        input.setSelectionRange(position, position + 1);
      }, 100);
    }
  }
  
  function showAllErrors(errors) {
    currentErrors = errors;
    currentErrorIndex = 0;
    showNextError();
  }
  
  function showNextError() {
    if (currentErrorIndex < currentErrors.length) {
      const error = currentErrors[currentErrorIndex];
      showErrorDialog(error.message, error.position);
    } else {
      // All errors handled, proceed with submission
      intentionalError = true;
      document.querySelector('form').submit();
    }
  }
  
  function closeErrorDialog() {
    document.getElementById('errorDialog').style.display = 'none';
    document.getElementById('input').focus();
  }
  
  function proceedWithError() {
    const checkbox = document.getElementById('intentionalCheck');
    if (!checkbox.checked) {
      alert('Please check the box if this is intentional, or click "Fix It" to correct the issue.');
      return;
    }
    
    closeErrorDialog();
    
    // Move to next error or submit if all handled
    currentErrorIndex++;
    if (currentErrorIndex < currentErrors.length) {
      showNextError();
    } else {
      // All errors marked as intentional, submit form
      intentionalError = true;
      const form = document.querySelector('form');
      const hiddenInput = document.createElement('input');
      hiddenInput.type = 'hidden';
      hiddenInput.name = 'intentional';
      hiddenInput.value = 'true';
      form.appendChild(hiddenInput);
      form.submit();
    }
  }
</script>

</body>
</html>`

// findAvailablePort checks for available port starting from given port
func findAvailablePort(startPort int) int {
	for port := startPort; port < startPort+100; port++ {
		ln, err := net.Listen("tcp", ":" + strconv.Itoa(port))
		if err == nil {
			ln.Close()
			return port
		}
	}
	return startPort // fallback
}

func main() {
	tmpl := template.Must(template.New("index").Parse(htmlTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{}

		// Track session
		sessionID := r.Header.Get("X-Forwarded-For")
		if sessionID == "" {
			sessionID = r.RemoteAddr
		}
		sessionMutex.Lock()
		activeSessions[sessionID] = true
		sessionMutex.Unlock()

		if r.Method == "POST" {
			input := r.FormValue("input")
			intentional := r.FormValue("intentional")
			if input != "" {
				// Skip validation if user marked as intentional
				var output string
				if intentional == "true" {
					output = processor.ProcessTextUnsafe(input)
				} else {
					output = processor.ProcessText(input)
				}
				
				if strings.HasPrefix(output, "ERROR:") {
					data.Error = output[7:] // Remove 'ERROR: ' prefix
					data.Input = input
				} else {
					data.Output = output
					data.Input = input
				}
			}
		}

		tmpl.Execute(w, data)
	})

	// Shutdown endpoint for browser close detection
	shutdownChan := make(chan bool, 1)
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"shutting down"}`))
		
		// Remove session
		sessionID := r.Header.Get("X-Forwarded-For")
		if sessionID == "" {
			sessionID = r.RemoteAddr
		}
		sessionMutex.Lock()
		delete(activeSessions, sessionID)
		sessionMutex.Unlock()
		
		// Shutdown when any browser closes
		go func() {
			time.Sleep(200 * time.Millisecond) // Give time for response
			select {
			case shutdownChan <- true:
			default:
			}
		}()
	})

	// Find available port
	port := findAvailablePort(8080)
	url := fmt.Sprintf("http://localhost:%d", port)
	fmt.Printf("🌐 Go Reloaded Web UI starting at %s\n", url)
	fmt.Println("💻 Press Ctrl+C to stop the server")
	
	// Auto-open browser after short delay
	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser(url)
	}()
	
	// Create server with graceful shutdown
	srv := &http.Server{Addr: ":" + strconv.Itoa(port)}
	
	// Start server in goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	
	// Wait for interrupt signal or browser close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	select {
	case <-c:
		fmt.Println("\n🛑 Manual shutdown requested...")
	case <-shutdownChan:
		fmt.Println("\n🛑 Browser closed, shutting down server...")
	}
	
	// Graceful shutdown
	fmt.Println("\n🛑 Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("✅ Server stopped")
}

// openBrowser opens the default browser to the given URL
func openBrowser(url string) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // linux
		cmd = "xdg-open"
		args = []string{url}
	}

	exec.Command(cmd, args...).Start()
}