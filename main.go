package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-korea/codelab/config"
	"github.com/golang-korea/codelab/handler"
)

const bind = ":1333"

func init() {
	config.Load()
}

func main() {
	// Register handlers
	http.HandleFunc("/", handler.RenderIndex)
	http.HandleFunc("/auth", handler.RenderLogin)
	http.HandleFunc("/auth/callback", handler.Authenticate)
	http.HandleFunc("/profile", handler.RenderProfile)

	// Start web server on 1333 port
	fmt.Printf("Server is listening on %s\n", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}