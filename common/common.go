package common

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var TestUsers = []User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
	{ID: "3", Name: "Charlie"},
}

func GetViteProxy() *httputil.ReverseProxy {
	// Define your Vite dev server URL (typically runs on port 5173)
	viteURL, err := url.Parse("http://localhost:5173")
	if err != nil {
		log.Fatal("Failed to parse Vite URL:", err)
	}

	// Create reverse proxy for Vite dev server
	viteProxy := httputil.NewSingleHostReverseProxy(viteURL)

	// Enhanced error handling for Vite proxy
	viteProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Vite proxy error: %v", err)
		http.Error(w, "Vite dev server unavailable", http.StatusBadGateway)
	}

	return viteProxy
}

// Check if Vite dev server is running
func IsViteServerRunning() bool {
	conn, err := net.Dial("tcp", "localhost:5173")
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func IsProduction() bool {
	return strings.Contains(strings.ToLower(os.Getenv("GO_ENV")), "prod")
}
