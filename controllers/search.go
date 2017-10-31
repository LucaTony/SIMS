package controllers

import (
    "fmt"
    "net/http"
    //"strings"
    _ "reflect"
    _ "bytes"

    "github.com/LucaTony/SIMS/models"
)

var mySearch string


//Testing POST
func (t *Todo) SearchPost() {
    fmt.Println("SearchPost")
    req := t.Ctx.Request()
    _ = req.ParseForm()

    mySearch = req.PostForm.Get("searchinput")
    fmt.Println(mySearch)

    //t.HTML(http.StatusOK)
    //t.Ctx.Template = "/test" //Post doesn't need a template

    t.Ctx.Redirect("/search", http.StatusFound)
}


//Testing GET
func (t *Todo) SearchGet() {
    fmt.Println("SearchGet")
    //t.HTML(http.StatusOK)
    todos := []*models.Todo{} // Create empty slice of struct pointers.
    //t.Ctx.Data["Data"] = S
    //t.Ctx.Redirect("/test", http.StatusFound) //don't redirect

    t.Ctx.DB.Order("created_at desc").Find(&todos)
    //var t.Ctx.Data["List"]
    teststring := []string{}

    for i := range todos {
        if todos[i].Body == mySearch {
            teststring = append(teststring, todos[i].Body)
            fmt.Println("added!!")
        }
    }

    //t.Ctx.Data["Data"] = []string{"This is", "Spartaa"}
    t.Ctx.Data["Data"] = teststring

    //fmt.Println(reflect.TypeOf(t.Ctx.Data["List"]))
    //for _, v := range todos {fmt.Printf("%v",v.Body)} // Iterate throuh objects, only value (since the _)

    //t.Ctx.Template = "index"
    t.Ctx.Template = "found"
    t.HTML(http.StatusOK)

}
