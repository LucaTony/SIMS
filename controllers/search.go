package controllers

import (
    "fmt"
    "net/http"
    "strings"
    _ "bytes"
    _ "reflect"

    "github.com/LucaTony/SIMS/models"
)

var mySearch string

type DataSend struct {
    Title       string
    Body        string
    Url        string
}


//Search POST
func (t *Todo) SearchPost() {
    fmt.Println("SearchPost")
    req := t.Ctx.Request()
    _ = req.ParseForm()

    mySearch = req.PostForm.Get("searchinput")
    fmt.Println(mySearch)

    //t.HTML(http.StatusOK)
    //t.Ctx.Template = "/test" //Post doesn't need a template

    t.Ctx.Redirect("/#search", http.StatusFound)
}


//Root GET
func (t *Todo) Home() {
    fmt.Println("SearchGet")
    todos := []*models.Todo{} // Create empty slice of struct pointers.
    //t.Ctx.Redirect("/test", http.StatusFound) //don't redirect

    t.Ctx.DB.Order("created_at desc").Find(&todos)
    //DataBody := []string{}


    var mySend = []DataSend {}

    for i := range todos {
        if (strings.Contains(strings.ToLower(todos[i].Body),  strings.ToLower(mySearch))) { //TODO: fix empty string
            //DataBody = append(DataBody, todos[i].Body)

            tempSend := DataSend{
                Title: todos[i].Title,
                Body: todos[i].Body,
                Url: todos[i].URL,
            }

            mySend = append(mySend, tempSend)
            fmt.Println("added!!")
        }
    }


    //fmt.Println(m["route"])

    //fukdis

    //t.Ctx.Data["Data"] = []string{"This is", "Spartaa"}

    t.Ctx.Data["Data"] = mySend
    //t.Ctx.Data["List"] = todos //Old one
    mySearch = ""

    //fmt.Println(reflect.TypeOf(t.Ctx.Data["List"]))
    //for _, v := range todos {fmt.Printf("%v",v.Body)} // Iterate throuh objects, only value (since the _)

    t.Ctx.Template = "index"
    t.HTML(http.StatusOK)

}
