package controllers

import (
    _ "bytes"
    "fmt"
    "net/http"
    _ "reflect"
    "sort"
    "strings"
    "time"

    "github.com/LucaTony/SIMS/models"
    "github.com/gernest/utron/controller"
)

var mySearch string // Text from the user search input field

// SearchSend for the Data to send to the HTML
type SearchSend struct {
    Title string
    Body  string
    Url   string
    Score int
    Found bool
}

//Search is a controller for a Search list
type Search struct {
    controller.BaseController
    Routes []string
}

type ByScore []SearchSend

// Functions for sorting the Search results by SearchSend.Score
func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

//Search gets the data from the searchinput field
func (t *Search) HousePost() {
    fmt.Println("SearchPost")
    req := t.Ctx.Request()
    _ = req.ParseForm()

    mySearch = req.PostForm.Get("searchinput")
    fmt.Println(mySearch)

    //t.Ctx.Template = "/test" //Post doesn't need a template
    t.Ctx.Redirect("/#search", http.StatusFound)
}

//Home is the root get that shows the webpage and the search results if the search function is used
//BUG(r): The searchinput should remain in the field when reloading
func (t *Search) Home() {
    fmt.Println("Home")

    //cal := &Calc{}
    //var cal *Calc
    //cal.CalcGet()
    t.CalcGet()
    fmt.Println("CalcGet: Done")
    t.FactGet()
    fmt.Println("FactGet: Done")

    t.Ctx.Template = "index"
    fmt.Println("Template: Done")
    t.HTML(http.StatusOK)
    fmt.Println("Status: Done")
}

//TODO: Refactor and understand
func (c Search) CalcGet(){
    questions := []*models.Calculator{}
    c.Ctx.DB.Order("id").Find(&questions)
    //fmt.Println(questions)

    //Get the question from the db.
    //Get the responding answer options to each question.
    c.Ctx.Data["CalcSend"] = questions
    //c.Ctx.Data["add"] = add(3,4)
    if (myScoreSet) {
        c.Ctx.Data["CalcResult"] = myScore
        c.Ctx.Data["CalcRecomm"] = myRecomm
        c.Ctx.Data["ResultLinkSend"] = myResultID
    }

    myScoreSet = false //Flush
    myRecomm = myRecomm[:0] //Flush

}
func (c Search) FactGet() {
    facts := []*models.FunFact{}
    c.Ctx.DB.Order("id").Find(&facts)
    c.Ctx.Data["Fact"] = facts

}

func (t *Search) AjaxSearch(){
    start := time.Now()
    mySearch := t.Ctx.Params["aj"]
    if len(mySearch) >= 3 {
        searches := []*models.Search{} // Create empty slice of struct pointers.
        t.Ctx.DB.Find(&searches)

        fmt.Println("AJAX: ", mySearch)
        //t.Ctx.Data["AjaxData"] = mySearch
        //t.Ctx.Commit()

        var mySend = []SearchSend{}

        if mySearch != "" {
            words := strings.Fields(mySearch)
            for i := range searches {
                score := 0
                for w := range words {
                    tempScore := strings.Count(strings.ToLower(searches[i].Body), strings.ToLower(words[w]))
                    if tempScore != 0 {
                        score += tempScore // add the amount one searched word occurs in a result
                    }
                }
                if score > 0 {
                    tempSend := SearchSend{ //TODO: Direct
                        Title: searches[i].Title,
                        Body:  searches[i].Body,
                        Url:   searches[i].URL,
                        Score: score,
                        Found: true,
                    }
                    mySend = append(mySend, tempSend)
                    //fmt.Println("added!!")
                }
            }
            if len(mySend) == 0 {
                tempSend := SearchSend{
                    Title: "Inga träffar",
                    Body:  mySearch,
                    Found: false,
                }
                mySend = append(mySend, tempSend)
            }

        }

        sort.Sort(ByScore(mySend))
        //t.Ctx.Data["Data"] = mySend
        //TODO: clean up this mess, especially direct without temp
        var myResults string
        for v := range mySend {
            myResults += `
            <div class="card text-left mt-2">
            <div class="card-header">
            `+ mySend[v].Title + `
            </div>
            <div class="card-body text-left">
            `
            if mySend[v].Found {
                myResults += `<p class="card-text">
                `+mySend[v].Body+ `
                </p><a target="_blank" href="` +mySend[v].Url+ `" class="btn btn-primary">Gå till hemsida</a>`
            } else {
                myResults += `<p class="card-text">Din sökning - <strong>`+mySend[v].Body+`</strong> - matchade inte något dokument</p>`
            }
            myResults += "</div></div>"
        }
        fmt.Fprintf(t.Ctx.Response(), myResults) //This responses like php's echo
        mySearch = "" // TODO: don't remove the searchfield text
    }
    //else {
        //fmt.Println("Not searching, too short")
    //}
        elapsed := time.Since(start)
        fmt.Println("It took: ", elapsed)
}

//NewSearch returns a new search list controller
func NewSearch() controller.Controller {
    return &Search{
        Routes: []string{
            "get;/;Home",
            "post;/search;HousePost",
            "get;/ajax/{aj};AjaxSearch",
        },
    }
}





