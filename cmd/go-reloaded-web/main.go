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
  --color-bg: #4a5568;
  --color-bg-panel: #ffffff;
  --color-text: #222;
  --color-accent: #5a5fff;
  --color-accent-gradient: linear-gradient(135deg, #4a4fdf, #8b4cdf);
  --color-border: #dcdfe6;
  --color-shadow: rgba(0, 0, 0, 0.05);
  --radius: 10px;
  --transition: 0.2s ease;
  --font-sans: "Inter", "Segoe UI", sans-serif;
  --font-mono: "Fira Code", monospace;
}

@media (prefers-color-scheme: dark) {
  :root {
    --color-bg: #11131a;
    --color-bg-panel: #1a1d26;
    --color-text: #e5e7eb;
    --color-accent: #7d7fff;
    --color-accent-gradient: linear-gradient(135deg, #7d7fff, #b57cff);
    --color-border: #2a2d3a;
    --color-shadow: rgba(0, 0, 0, 0.4);
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

.transform-btn {
  margin-top: 16px;
  background: var(--color-accent-gradient);
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 12px;
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
  margin-top: 8px;
  background: #dc3545;
  color: #fff;
  border: none;
  border-radius: var(--radius);
  padding: 8px;
  font-weight: 500;
  font-size: 0.9rem;
  cursor: pointer;
  transition: var(--transition);
}

.clear-btn:hover {
  background: #c82333;
  transform: translateY(-1px);
}

.how-to {
  margin-top: 12px;
  padding: 12px;
  background: rgba(90, 95, 255, 0.1);
  border-radius: var(--radius);
  font-size: 0.85rem;
  line-height: 1.4;
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
  <div class="subtitle">Text Transformation Suite</div>
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
      <button type="submit" class="transform-btn">Transform Text</button>
      <button type="button" class="clear-btn" onclick="clearText()">Clear</button>
    </form>
    <div class="how-to">
      <h4>How to Use</h4>
      • Click buttons to insert transformations<br>
      • Single Enter: New line<br>
      • Double Enter: Transform text<br>
      • Use Transform button or double Enter
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
  <span>Text Processing Engine v1.1</span>
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
  
  function clearText() {
    document.getElementById('input').value = '';
    document.getElementById('output').value = '';
    document.getElementById('input').focus();
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
    }
  });
  
  window.addEventListener('beforeunload', function() {
    navigator.sendBeacon('/shutdown');
  });
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
			if input != "" {
				output := processor.ProcessText(input)
				data.Output = output
				data.Input = input
			}
		}

		tmpl.Execute(w, data)
	})

	// Shutdown endpoint for browser close detection
	shutdownChan := make(chan bool, 1)
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		
		// Remove session
		sessionID := r.Header.Get("X-Forwarded-For")
		if sessionID == "" {
			sessionID = r.RemoteAddr
		}
		sessionMutex.Lock()
		delete(activeSessions, sessionID)
		sessionCount := len(activeSessions)
		sessionMutex.Unlock()
		
		// Shutdown when any browser closes
		select {
		case shutdownChan <- true:
		default:
		}
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