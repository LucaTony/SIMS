//Package controllers handles the different routes and is the interface between
// the model and the view
package controllers

import (
    //"fmt"
    //_ "github.com/LucaTony/SIMS/models"
    "github.com/gernest/utron/controller"
    "github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

//Search is a controller for a Search list
type Search struct {
    controller.BaseController
    Routes []string
}

//NewSims returns a new sims list controller
func NewSims() controller.Controller {
    return &Search{
        Routes: []string{
            "get;/;Home",

            "post;/search;HousePost",
            "get;/search;HouseGet",

            "post;/create;Create",
            "get;/delete/{id};Delete",
            "get;/scraper;ScraperGet",
            "post;/scraper;ScraperPost",
            "get;/calculator;CalcGet",
            "post;/calculator;CalcPost",
        },
    }
}
