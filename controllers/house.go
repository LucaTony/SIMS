package controllers

import (
	"fmt"
	//"net/http"
)

// Show Test
func (t *Todo) Show() {
    fmt.Println("showing Model")
    //t.HTML(http.StatusOK)
    //t.Ctx.Redirect("/test", http.StatusFound)
}
