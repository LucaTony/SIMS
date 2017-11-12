package controllers

import(
    "fmt"
    "net/http"
    "github.com/anaskhan96/soup"
    "os"
)

//UpdateGet is a scraper to update the database
func (t *Todo) UpdateGet() {
    fmt.Println("UpdateGet")
    //resp, err := soup.Get("https://xkcd.com")
    resp, err := soup.Get("http://www.energimyndigheten.se/tester")
    if err != nil {
        os.Exit(1)
    }
    doc := soup.HTMLParse(resp)

    links := doc.Find("div", "class", "intro").Find("p")
    fmt.Println(links.Text())

    //for _, link := range links {
        //fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    //}

    t.Ctx.Template = "_index"
    t.HTML(http.StatusOK)
}

