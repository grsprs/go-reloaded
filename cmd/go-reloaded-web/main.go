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
	"syscall"
	"time"

	"go-reloaded/internal/processor"
)

type PageData struct {
	Input  string
	Output string
	Error  string
}

const htmlTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Go Reloaded</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 0; padding: 20px; background: #f0f2f5; }
        .container { max-width: 1400px; margin: 0 auto; background: white; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .header { background: #4a90e2; color: white; padding: 20px; text-align: center; }
        .header h1 { margin: 0; font-size: 28px; }
        .main { display: grid; grid-template-columns: 250px 1fr 250px; gap: 20px; padding: 20px; min-height: 600px; }
        .sidebar { background: #f8f9fa; padding: 15px; border-radius: 8px; border: 1px solid #e9ecef; }
        .sidebar h3 { margin: 0 0 15px 0; color: #495057; font-size: 16px; }
        .btn-transform { width: 100%; background: #28a745; color: white; border: none; padding: 10px; margin: 5px 0; border-radius: 5px; cursor: pointer; font-size: 13px; }
        .btn-transform:hover { background: #218838; }
        .example { font-size: 11px; color: #6c757d; margin: 5px 0 15px 0; font-family: monospace; }
        .center { background: white; padding: 20px; border: 1px solid #e9ecef; border-radius: 8px; }
        .form-group { margin-bottom: 20px; }
        .form-group label { display: block; margin-bottom: 8px; font-weight: bold; color: #495057; }
        .textarea { width: 100%; height: 250px; padding: 15px; border: 2px solid #ced4da; border-radius: 5px; font-family: monospace; font-size: 14px; resize: vertical; }
        .textarea:focus { outline: none; border-color: #4a90e2; }
        .btn-main { background: #4a90e2; color: white; border: none; padding: 15px 30px; border-radius: 5px; font-size: 16px; cursor: pointer; width: 100%; }
        .btn-main:hover { background: #357abd; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Go Reloaded - Text Transformation Tool</h1>
        </div>
        <div class="main">
            <div class="sidebar">
                <h3>Numbers</h3>
                <button class="btn-transform" onclick="insertText('42 (hex)')">Hex to Decimal</button>
                <div class="example">42 (hex) → 66</div>
                <button class="btn-transform" onclick="insertText('1010 (bin)')">Binary to Decimal</button>
                <div class="example">1010 (bin) → 10</div>
                <h3>Case</h3>
                <button class="btn-transform" onclick="insertText('word (up)')">Uppercase</button>
                <div class="example">word (up) → WORD</div>
                <button class="btn-transform" onclick="insertText('word (low)')">Lowercase</button>
                <div class="example">word (low) → word</div>
            </div>
            <div class="center">
                <form method="POST">
                    <div class="form-group">
                        <label>Input Text</label>
                        <textarea name="input" id="input" class="textarea" placeholder="Enter your text here...">{{.Input}}</textarea>
                    </div>
                    <div class="form-group">
                        <label>Output</label>
                        <textarea id="output" class="textarea" readonly placeholder="Transformed text will appear here...">{{.Output}}</textarea>
                    </div>
                    <button type="submit" class="btn-main">Transform Text</button>
                </form>
            </div>
            <div class="sidebar">
                <h3>Format</h3>
                <button class="btn-transform" onclick="insertText('word ,')">Fix Punctuation</button>
                <div class="example">word , → word,</div>
                <button class="btn-transform" onclick="insertText(&quot;' text '&quot;)">Fix Quotes</button>
                <div class="example">' text ' → 'text'</div>
                <h3>Grammar</h3>
                <button class="btn-transform" onclick="insertText('a apple')">Fix Articles</button>
                <div class="example">a apple → an apple</div>
            </div>
        </div>
    </div>
    <script>
        function insertText(text) {
            const input = document.getElementById('input');
            input.value += text + ' ';
            input.focus();
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