package controllers

import (
    "fmt"
    "time"
    "net/http"
    "os"
    _ "reflect"
    "strings"
    "strconv"
    "github.com/anaskhan96/soup"
    "github.com/LucaTony/SIMS/models"
    "github.com/gernest/utron/controller"
    "github.com/antzucaro/matchr"
)

//TODO: TrimLeft / 

var updated bool = false
var updated_urls []string

//Scraper is a controller for a Search list
type Scraper struct {
    controller.BaseController
    Routes []string
}

//ScraperGet is a scraper to update the database
func (t *Scraper) ScraperGet() {
    fmt.Println("ScraperGet")

    //Check if Post got called
    if updated {
        fmt.Println("updated..")
        t.Ctx.Data["Urls"] = updated_urls
        updated_urls = updated_urls[:0] //Flush
        updated = false // Flush
    }

    t.Ctx.Template = "scraper"
    t.HTML(http.StatusOK)
}

//CheckDue checks if a update is necessary by looking at the
//time since last update and the amount of change
func CheckDue (existing string, recent string, updated time.Time) bool {
    t := time.Now()
    var updateInterval time.Duration = 24*time.Hour
    var due bool

    // Check if entry is empty
    if existing == "" {
        return true
    }
    //Check if entry is identical
    if existing == recent {
        return false
    }

    //Older than 1 day
    if t.Sub(updated) > updateInterval {
        //Use the DamerauLevenshtein algortithm to calculate the distance/
        //cost for two strings. Normalize with length
        var change = float32(matchr.DamerauLevenshtein(existing,recent))/float32(len(existing))
        if change > 0.8 {
            due = false
            fmt.Println("Text changed too much, not updating..")
        } else { due = true }
    } else { due = false }

    return due
}

//ScraperPost is the Post function for the scaper
func (t *Scraper) ScraperPost() {
    fmt.Println("ScraperPost")
    updated = true // Flag for Get

    searches := []*models.Search{} // Create empty slice of struct pointers.

    //scraperID := t.Ctx.Params["id"] // Id which should be updated
    req := t.Ctx.Request()
    _ = req.ParseForm()

    //scraperID := req.ParseForm["scraperID"]
    scraperID := req.PostForm.Get("scraperID")

    if scraperID == "all" {
        t.Ctx.DB.Order("id").Find(&searches)
    } else {
        ID, _ := strconv.Atoi(scraperID)
        t.Ctx.DB.Order("id").First(&searches, ID)
    }

    swe, err := time.LoadLocation("Europe/Stockholm")
    if err != nil {
        fmt.Println("err: ", err.Error())
    }


    for i := range searches {
        if (searches[i].Dom != "") && (searches[i].DomTitle != "") {
            fmt.Println("Proccessing id: ", searches[i].ID)
            resp, err := soup.Get(searches[i].URL)
            if err != nil {
                fmt.Println("Couldn't get the URL")
                os.Exit(1)
            }
            doc := soup.HTMLParse(resp)

            //Get the Body text
            find := strings.Split(searches[i].Dom, ".")
            find0 := strings.Split(find[0],",") // div class intro
            find1 := strings.Split(find[1],"|") // p 3
            tagIndex, err := strconv.Atoi(find1[1]) // 3
            sitesBody := doc.Find(find0...).FindAll(find1[0])[tagIndex]

            //Get the Title text
            find = strings.Split(searches[i].DomTitle, ".")
            find0 = strings.Split(find[0],",")
            find1 = strings.Split(find[1],"|")
            tagIndex, err = strconv.Atoi(find1[1])
            sitesTitle := doc.Find(find0...).FindAll(find1[0])[tagIndex]

            //for f := range find0 { fmt.Println("The first tags are: ", find0[f]) }

            //Check if update is due and update the DB entry
            if CheckDue(searches[i].Body, sitesBody.Text(), searches[i].UpdatedAt.In(swe)) {
                fmt.Println("Updating: ", searches[i].URL)
                //TODO: Don't misuse i as id !!
                //t.Ctx.DB.Order("id").First(&searches, ID)
                t.Ctx.DB.Model(&searches).Where("id="+ strconv.Itoa(searches[i].ID)).Updates(map[string]interface{}{
                    "body" : sitesBody.Text(),
                    "title" : sitesTitle.Text(),
                })
                fmt.Println("New Title: ", sitesTitle.Text())
                fmt.Println("New Body: ", sitesBody.Text())
                updated_urls = append(updated_urls , searches[i].URL) // Show updated urls on GET
            } else {
                fmt.Println("Not Updated: ", searches[i].URL)
            }
        }
    }
    t.Ctx.Redirect("/scraper", http.StatusFound)
}

//NewScraper returns a new scraper list controller
func NewScraper() controller.Controller {
    return &Scraper{
        Routes: []string{
            "get;/scraper;ScraperGet",
            "post;/scraper;ScraperPost",
        },
    }
}
