package controllers

import (
	"fmt"
	"net/http"
)

//NotFound Custom 404
func (t *Todo) Show() {
    fmt.Println("showing Model")
    t.HTML(http.StatusOK)
    t.Ctx.Redirect("/", http.StatusFound)
}