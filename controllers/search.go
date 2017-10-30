package controllers

import (
	"fmt"
    "net/http"
    //"strings"
    //"reflect"
)

var S string


//Testing POST
func (t *Todo) SearchPost() {
    fmt.Println("SearchPost")
	req := t.Ctx.Request()
    _ = req.ParseForm()

    S = req.PostForm.Get("searchinput")
    fmt.Println(S)

    //t.HTML(http.StatusOK)
    //t.Ctx.Template = "/test" //Post doesn't need a template

    t.Ctx.Redirect("/search", http.StatusFound)
}


//Testing GET
func (t *Todo) SearchGet() {
    fmt.Println("SearchGet")
    //t.HTML(http.StatusOK)
	t.Ctx.Data["Data"] = S
    t.Ctx.Template = "found"
    //t.Ctx.Redirect("/test", http.StatusFound) //don't redirect
}
