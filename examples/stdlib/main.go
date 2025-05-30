package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"spa/common"
	"spa/ui"
)

func main() {
	// Create main router
	mux := http.NewServeMux()

	// API routes - handled by Go backend
	mux.HandleFunc("/api/users", handleUsers)
	mux.HandleFunc("/api/users/{id}", handleUserByID)
	mux.HandleFunc("/health", handleHealth)

	if common.IsProduction() {
		hfs := http.FileServer(http.FS(ui.DistDirFS))
		// Serve static assets from the ui/dist directory
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
				// If the file does not exist, serve index.html for SPA routing
				if !doesFileExist(ui.DistDirFS, r.URL.Path) {
					http.ServeFileFS(w, r, ui.DistDirFS, "index.html")
					return
				}
			}
			hfs.ServeHTTP(w, r)
		})
	} else {
		// Check if Vite dev server is running
		if !common.IsViteServerRunning() {
			log.Fatalln("Vite dev server is not running. Please start it with 'pnpm dev' from the ui/ directory")
			return
		}

		// Proxy to Vite dev server for frontend assets and routes in non-production
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			viteProxy := common.GetViteProxy()

			// Proxy to Vite for frontend assets and routes
			viteProxy.ServeHTTP(w, r)
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Println("📡 API endpoints: /api/*")
	fmt.Println("💚 Health check: /health")
	if common.IsProduction() {
		fmt.Println("🌐 Serving static files from ui/dist directory")
	} else {
		fmt.Println("🔄 Proxying to Vite dev server at http://localhost:5173")
	}

	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// Example API handler
func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(common.TestUsers)
		return
	case http.MethodPost:
		fmt.Fprint(w, `{"message": "User created successfully"}`)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		var user common.User
		for _, u := range common.TestUsers {
			if u.ID == id {
				user = u
				break
			}
		}
		json.NewEncoder(w).Encode(user)
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status": "healthy", "service": "go-backend"}`)
}

func doesFileExist(f fs.FS, path string) bool {
	_, err := fs.Stat(f, path)
	return err == nil
}
