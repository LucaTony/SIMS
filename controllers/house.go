package controllers

import (
	"fmt"
    "net/http"
    //"strings"
    //"reflect"
)

var X string

// Show Test
func (t *Todo) TestPost() {
    fmt.Println("TestPost")
	req := t.Ctx.Request()
    _ = req.ParseForm()

    X = req.PostForm.Get("textinput")
    fmt.Println(X)

    //t.HTML(http.StatusOK)
    //t.Ctx.Template = "/test" //Post doesn't need a template
	t.Ctx.Redirect("/test", http.StatusFound)
}

func (t *Todo) TestGet() {
    fmt.Println("TestGet")
    //t.HTML(http.StatusOK)
	t.Ctx.Data["Data"] = X
    t.Ctx.Template = "test"
    //t.Ctx.Redirect("/test", http.StatusFound) //don't redirect
}
