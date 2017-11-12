//Package controllers handles the different routes and is the interface between
// the model and the view
package controllers

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/LucaTony/SIMS/models"
    "github.com/gernest/utron/controller"
    "github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

//Todo is a controller for Todo list
type Todo struct {
    controller.BaseController
    Routes []string
}

//Create creates a todo  item
func (t *Todo) Create() {
    fmt.Println("Creating..")
    todo := &models.Todo{}
    req := t.Ctx.Request()
    _ = req.ParseForm()
    if err := decoder.Decode(todo, req.PostForm); err != nil {
        t.Ctx.Data["Message"] = err.Error()
        t.Ctx.Template = "error"
        t.HTML(http.StatusInternalServerError)
        return
    }

    t.Ctx.DB.Create(todo)
    t.Ctx.Redirect("/", http.StatusFound)
}

//Delete deletes a todo item
func (t *Todo) Delete() {
    todoID := t.Ctx.Params["id"]
    ID, err := strconv.Atoi(todoID)
    if err != nil {
        t.Ctx.Data["Message"] = err.Error()
        t.Ctx.Template = "error"
        t.HTML(http.StatusInternalServerError)
        return
    }
    t.Ctx.DB.Delete(&models.Todo{ID: ID})
    t.Ctx.Redirect("/", http.StatusFound)
}

//Testing is a test function
func (t *Todo) Testing() {
    fmt.Println("Testing")
}


//NewTodo returns a new  todo list controller
func NewTodo() controller.Controller {
    return &Todo{
        Routes: []string{
            "get;/;Home",

            "post;/search;SearchPost",
            "get;/search;SearchGet",

            "post;/create;Create",
            "get;/delete/{id};Delete",
            "get;/update;UpdateGet",
        },
    }
}
