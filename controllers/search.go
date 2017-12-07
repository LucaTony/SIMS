package controllers

import (
	_ "bytes"
	"fmt"
	"net/http"
	_ "reflect"
	"sort"
	"strings"

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
	searches := []*models.Search{} // Create empty slice of struct pointers.

	t.Ctx.DB.Order("created_at desc").Find(&searches)

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
				Title: "Error",
				Body:  mySearch,
				Found: false,
			}
			mySend = append(mySend, tempSend)
		}

	}

	sort.Sort(ByScore(mySend))
	t.Ctx.Data["Data"] = mySend
	mySearch = "" // TODO: don't remove the searchfield text
	//for _, v := range searches {fmt.Printf("%v",v.Body)} // Iterate throuh objects, only value (since the _)

	//cal := &Calc{}
	//var cal *Calc
	//cal.CalcGet()
	t.CalcGet()
	t.FactGet()

	t.Ctx.Template = "index"
	t.HTML(http.StatusOK)
}

//TODO: Refactor and understand
<<<<<<< HEAD
func (c Search) CalcGet() {
	questions := []*models.Calculator{}
	c.Ctx.DB.Order("id").Find(&questions)
	//fmt.Println(questions)
	mySend := []CalcSend{}

	//Get the question from the db.
	//Get the responding answer options to each question.
	for q := range questions {
		tempSend := CalcSend{
			ID:       questions[q].ID,
			Question: questions[q].Question,
			Option01: questions[q].Option01,
			Option02: questions[q].Option02,
			Option03: questions[q].Option03,
			Option04: questions[q].Option04,
			Option05: questions[q].Option05,
			Option06: questions[q].Option06,
		}
		mySend = append(mySend, tempSend)
		//fmt.Println(mySend)
	}
	c.Ctx.Data["CalcSend"] = mySend
	//c.Ctx.Data["add"] = add(3,4)
	if myScoreSet {
		c.Ctx.Data["CalcResult"] = myScore
		c.Ctx.Data["CalcRecomm"] = myRecomm
		c.Ctx.Data["ResultLinkSend"] = myResultID
	}

	myScoreSet = false      //Flush
	myRecomm = myRecomm[:0] //Flush
=======
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
>>>>>>> dfd75cfa71185a4f227dc93d9c9352a304562b4a

}
func (c Search) FactGet() {
	facts := []*models.FunFact{}
	c.Ctx.DB.Order("id").Find(&facts)
	c.Ctx.Data["Fact"] = facts

}

//NewSearch returns a new search list controller
func NewSearch() controller.Controller {
	return &Search{
		Routes: []string{
			"get;/;Home",
			"post;/search;HousePost",
		},
	}
}
