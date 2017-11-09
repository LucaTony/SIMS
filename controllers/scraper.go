package controllers

import(
    "fmt"
    "net/http"
    "github.com/anaskhan96/soup"
    "os"
)



func (t *Todo) UpdateGet() {
    fmt.Println("UpdateGet")
    resp, err := soup.Get("https://xkcd.com")
    if err != nil {
        os.Exit(1)
    }
    doc := soup.HTMLParse(resp)
    links := doc.Find("div", "id", "comicLinks").FindAll("a")
    for _, link := range links {
        fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    }

    t.Ctx.Template = "_index"
    t.HTML(http.StatusOK)
}

