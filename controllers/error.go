package controllers

import (
	"net/http"
)

//NotFound Custom 404
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The page was not found!"))
}
