package controllers

import(
    "fmt"
    "net/http"
    "github.com/anaskhan96/soup"
    "os"
    "github.com/LucaTony/SIMS/models"
    _ "reflect"
    "strings"
    "strconv"
)

//UpdateGet is a scraper to update the database
func (t *Todo) UpdateGet() {
    fmt.Println("UpdateGet")
    todos := []*models.Todo{} // Create empty slice of struct pointers.
    //todo := &models.Todo{}

    t.Ctx.DB.Order("id").Find(&todos)
    for i := range todos {
        if todos[i].Dom != "" {
            resp, err := soup.Get(todos[i].URL)
            if err != nil {
                os.Exit(1)
            }
            doc := soup.HTMLParse(resp)

            //fmt.Println(reflect.TypeOf(todos[i].Dom))
            find := strings.Split(todos[i].Dom, ".")
            find0 := strings.Split(find[0],",")
            find1 := find[1]
            sites := doc.Find(find0...).Find(find1)

            fmt.Println("parent-tag: ", find0)
            fmt.Println("tag: ", find1)
            fmt.Println(sites.Text())

            t.Ctx.DB.Model(&todos).Where("id="+ strconv.Itoa(i+1)).Select("body").Updates(map[string]interface{}{"body" : sites.Text()})
    //for _, link := range links {
    //fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
    //}
    //}
        }
    }

    t.Ctx.Template = "testing"
    t.HTML(http.StatusOK)
}

