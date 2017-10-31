package controllers

import (
    "fmt"
    "net/http"
    //"strings"
    //"reflect"

    "github.com/LucaTony/SIMS/models"
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

func remove(s []int, i int) []int {
    s[len(s)-1], s[i] = s[i], s[len(s)-1]
    return s[:len(s)-1]
}


//Testing GET
func (t *Todo) SearchGet() {
    fmt.Println("SearchGet")
    //t.HTML(http.StatusOK)
    todos := []*models.Todo{}
    //t.Ctx.Data["Data"] = S

    //t.Ctx.Redirect("/test", http.StatusFound) //don't redirect

    t.Ctx.DB.Order("created_at desc").Find(&todos)

    for i := range todos {
        if todos[i].Body != "test" {
            todos = append(todos[:i], todos[i+1:]...)
            fmt.Println("deleted!!")
        }
    }

    t.Ctx.Data["List"] = todos
    //fmt.Println(reflect.TypeOf(todos))
    //for _, v := range todos {fmt.Printf("%v",v.Body)} // Iterate throuh objects, only value (since the _)


    //t.Ctx.Template = "index"
    t.Ctx.Template = "found"
    t.HTML(http.StatusOK)


}
