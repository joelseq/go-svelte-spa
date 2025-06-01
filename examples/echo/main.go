package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"spa/common"
	"spa/ui"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodPatch},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderContentType},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/health", handleHealth)

	// API routes
	api := e.Group("/api")
	api.GET("/users", handleUsers)
	api.GET("/users/:id", handleUserByID)

	if common.IsProduction() {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			HTML5:      true,
			Root:       "dist",
			Filesystem: http.FS(ui.DistDir),
		}))
	} else {
		if !common.IsViteServerRunning() {
			log.Fatalln("Vite dev server is not running. Please start it with 'pnpm dev'")
		}
		// Static assets and frontend routes
		e.Any("/*", staticAssetHandler)
	}

	fmt.Println("📡 API endpoints: /api/*")
	fmt.Println("💚 Health check: /health")
	if common.IsProduction() {
		fmt.Println("🌐 Serving static files from ui/dist directory")
	} else {
		fmt.Println("🔄 Proxying to Vite dev server at http://localhost:5173")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func handleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "go-backend",
	})
}

func handleUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, common.TestUsers)
}

func handleUserByID(c echo.Context) error {
	id := c.Param("id")
	for _, user := range common.TestUsers {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
}

func staticAssetHandler(c echo.Context) error {
	viteProxy := common.GetViteProxy()
	// Proxy to Vite for frontend assets and routes
	viteProxy.ServeHTTP(c.Response().Writer, c.Request())
	return nil
}
