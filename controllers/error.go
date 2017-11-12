package controllers

import (
	"net/http"
)

//NotFound is the Custom 404 handler
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The page was not found!"))
}
