package controllers

import (
    "fmt"
    "net/http"
    "strings"
    _ "bytes"
    _ "reflect"
    "sort"
    "github.com/LucaTony/SIMS/models"
)

var mySearch string // Text from the user search input field

// Datatype for the Data to send to the HTML
type DataSend struct {
    Title string
    Body  string
    Url   string
    Score int
}

type ByScore []DataSend

// Functions for sorting the Search results by DataSend.Score
func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }


//Search POST gets the data from the searchinput field
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


//Root GET shows the webpage and the search results if the search function is used
//BUG(l): Should convert to lowercase
func (t *Todo) Home() {
    fmt.Println("SearchGet")
    todos := []*models.Todo{} // Create empty slice of struct pointers.

    t.Ctx.DB.Order("created_at desc").Find(&todos)

    var mySend = []DataSend {}

    if (mySearch != ""){
        words := strings.Fields(mySearch)
        for i := range todos {
            score := 0
            for w := range words {
                tempScore := strings.Count(strings.ToLower(todos[i].Body), strings.ToLower(words[w]))
                if (tempScore != 0) {
                    score += tempScore // add the amount one searched word occurs in a result
                }
            }
            if (score > 0) {
                tempSend := DataSend{ //TODO: Direct
                    Title: todos[i].Title,
                    Body:  todos[i].Body,
                    Url:   todos[i].URL,
                    Score: score,
                }
                mySend = append(mySend, tempSend)
                //fmt.Println("added!!")
            }
            //if (strings.Contains(strings.ToLower(w), strings.ToLower(todos[i].Body))) { // a searched word 
        }
    }

    //t.Ctx.Data["Data"] = []string{"This is", "Spartaa"}
    sort.Sort(ByScore(mySend))
    t.Ctx.Data["Data"] = mySend
    //t.Ctx.Data["List"] = todos //Old one
    mySearch = "" // TODO: don't remove the searchfield text

    //fmt.Println(reflect.TypeOf(t.Ctx.Data["List"]))
    //for _, v := range todos {fmt.Printf("%v",v.Body)} // Iterate throuh objects, only value (since the _)

    t.Ctx.Template = "index"
    t.HTML(http.StatusOK)

}
