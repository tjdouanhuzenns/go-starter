package code

import (
	"bytes"
	"text/template"
)

// routerTemplate defines the base router setup using chi router
const routerTemplate = `package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter initializes and returns a new chi router with default middlewares.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Built-in middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	return r
}
`

// routerMainTemplate defines the main server entry point
const routerMainTemplate = `package main

import (
	"fmt"
	"net/http"

	"{{.ModuleName}}/internal/router"
	"{{.ModuleName}}/pkg/config"
)

func main() {
	cfg := config.NewConfig()

	r := router.NewRouter()

	addr := fmt.Sprintf(":%d", cfg.App.Port)
	fmt.Printf("Server is running on %s\\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		panic(err)
	}
}
`

// RouterTemplateData holds data required to render the router main template.
type RouterTemplateData struct {
	ModuleName string
}

// RouterCodeGenerate generates the base router file content.
// It returns the rendered router code as a string.
func RouterCodeGenerate() (string, error) {
	tmpl, err := template.New("router").Parse(routerTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse router template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, nil); err != nil {
		return "", fmt.Errorf("failed to execute router template: %w", err)
	}

	return buf.String(), nil
}

// MainEntryCodeGenerate generates the main.go entry point content
// using the provided module name.
func MainEntryCodeGenerate(moduleName string) (string, error) {
	tmpl, err := template.New("main").Parse(routerMainTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse main entry template: %w", err)
	}

	data := RouterTemplateData{
		ModuleName: moduleName,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute main entry template: %w", err)
	}

	return buf.String(), nil
}
