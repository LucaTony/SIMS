package controllers

import (
    "fmt"
    "time"
    "net/http"
    "github.com/anaskhan96/soup"
    "os"
    "github.com/LucaTony/SIMS/models"
    "github.com/antzucaro/matchr"
    _ "reflect"
    "strings"
    "strconv"
)

//TODO: TrimLeft / 

var updated bool = false
var updated_urls []string

//ScraperGet is a scraper to update the database
func (t *Search) ScraperGet() {
    fmt.Println("ScraperGet")

    //Check if Post got called
    if updated {
        fmt.Println("updated..")
        t.Ctx.Data["Urls"] = updated_urls
        updated = false
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
func (t *Search) ScraperPost() {
    updated = true
    searches := []*models.Search{} // Create empty slice of struct pointers.
    //todo := &models.Todo{}
    fmt.Println("ScraperPost")
    t.Ctx.DB.Order("id").Find(&searches)

    swe, err := time.LoadLocation("Europe/Stockholm")
    if err != nil {
        fmt.Println("err: ", err.Error())
    }

    //updated_urls = append(updated_urls , "url1")
    for i := range searches {
        if searches[i].Dom != "" {
            resp, err := soup.Get(searches[i].URL)
            if err != nil {
                os.Exit(1)
            }
            doc := soup.HTMLParse(resp)

            //Get the Body text
            find := strings.Split(searches[i].Dom, ".")
            find0 := strings.Split(find[0],",") // div class intro
            find1 := strings.Split(find[1],"|") // p 3
            tagIndex, err := strconv.Atoi(find1[1])
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
                t.Ctx.DB.Model(&searches).Where("id="+ strconv.Itoa(i+1)).Updates(map[string]interface{}{
                    "body" : sitesBody.Text(),
                    "title" : sitesTitle.Text(),
                })
                //fmt.Println(sitesBody.Text())
                //fmt.Println(sitesBody.Text())
                updated_urls = append(updated_urls , searches[i].URL)
            } else {
                fmt.Println("Not Updating: ", searches[i].URL)
            }
        }
    }
    t.Ctx.Redirect("/scraper", http.StatusFound)
}


