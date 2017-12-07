package controllers

import (
    "fmt"
    "net/http"
    "github.com/LucaTony/SIMS/models"
    "github.com/gernest/utron/controller"
    "reflect"
    "strconv"
    "strings"
)

type CalcSend struct {
    ID          int
    Question    string
    Option01    string
    Option02    string
    Option03    string
    Option04    string
    Option05    string
    Option06    string
}

//Calc is a controller for a Calculator list
type Calc struct {
    controller.BaseController
    Routes []string
}
//Result is a controller for a Calculator list
type Result struct {
    controller.BaseController
    Routes []string
}

var myScore float32             //Score in percentage
var myScoreSet bool = false     //True if test was taken
var myRecomm = []string{}       //Show recomm on startpage
var myResultID int              //id, primary key for db for unique link
var myResultRecomm string       //Store Recomm for future access

func getField(c *models.Calculator, field string) float32 {
    r := reflect.ValueOf(c)
    f := reflect.Indirect(r).FieldByName(field)
    return float32(f.Float())
}

//CalcPost function to get questions and options from db
// and then display it on the website.
func (c *Calc) CalcPost() {
    fmt.Println("CalcPost")
    req := c.Ctx.Request()
    _ = req.ParseForm()
    //fmt.Println(req.PostForm)

    questions := []*models.Calculator{}
    c.Ctx.DB.Order("id").Find(&questions)

    myScore = 0
    myIndex := 0        //usually id-1
    recommInd := 0      //Index for recomm values
    //TODO: Refactor DBlookup
    // rq: Question1, o: [2]
    //Iteration order is like in db i.e. id
    for q := range questions {
        tempScore := getField(questions[q], "Points0" + req.PostForm["Question"+ strconv.Itoa(myIndex+1)][0] )
        if (tempScore == 0) {
            myRecomm = append(myRecomm, questions[q].Recomm)
            recommInd++
            myResultRecomm += strconv.Itoa(myIndex+1)
            myResultRecomm += (" ") //Delimiter hack
        }
        myScore += tempScore
        myIndex++
    }

    myScore *= 10
    myScoreSet = true
    fmt.Println("Your score is : ", myScore)

    c.ResultSave()
    c.Ctx.Redirect("/#calculator", http.StatusFound)
}

func (c *Calc) ResultSave() {
    result := &models.Result{
        Total: myScore,
        Recomm: myResultRecomm,
    }
    c.Ctx.DB.Create(result)
    myResultID = result.ID
    c.Ctx.Data["ResultLinkSend"] = myResultID
    fmt.Println("Your ID for the result is: ", result.ID)
}

func (c *Calc) ResultGet() {
    resultID := c.Ctx.Params["id"]
    ID, _ := strconv.Atoi(resultID) // int not needed
    //myFound := c.Ctx.DB.Select("total").Find(&models.Result{ID: ID})
    results := []*models.Result{}
    c.Ctx.DB.First(&results, ID)
    //fmt.Println(results[0])
    c.Ctx.Data["ResultSend"] = results[0]
    myTempRecommId := strings.Fields(results[0].Recomm) //get recomm ids
    myTempRecomm := []string{}

    calculators := []*models.Calculator{}
    for i := range myTempRecommId {
        tempid, err := strconv.Atoi(myTempRecommId[i])
        if err != nil {
            fmt.Println("RecommId couldn't be Atoi")
        }
        c.Ctx.DB.First(&calculators, tempid)
        myTempRecomm = append(myTempRecomm, calculators[0].Recomm)
        fmt.Println(calculators[0].Recomm)
    }

    c.Ctx.Data["SavedRecomm"] = myTempRecomm

    c.Ctx.Template = "result"
    c.HTML(http.StatusOK)
}


//NewCalc returns a new sims list controller
func NewCalc() controller.Controller {
    return &Calc{
        Routes: []string{
            "get;/calculator;CalcGet",
            "post;/calculator;CalcPost",
            "get;/save;ResultSave",
            "get;/result/{id};ResultGet",
            //"get;/;CalcPost",
        },
    }
}
