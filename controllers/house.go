package controllers

import (
	"fmt"
    "net/http"
    //"strings"
    //"reflect"
)

//var x

// Show Test
func (t *Todo) Test2() {
    //fmt.Println("Test2")
	req := t.Ctx.Request()
    _ = req.ParseForm()

    x := req.PostForm.Get("textinput")
    //fmt.Println(req)
    fmt.Println(x)

    //t.HTML(http.StatusOK)
    t.Ctx.Template = "/test"
	t.Ctx.Redirect("/test", http.StatusFound)
}

