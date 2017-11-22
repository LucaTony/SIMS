package models

import (
    "time"
)

//Person represent an item of search list
type Search struct {
    ID          int         `schema:"-"`
    Title       string      `schema:"title"`
    Body        string      `schema:"body"`
    URL	        string      `schema:"url"`
    Dom         string      `schema:"dom"`
    DomTitle    string      `schema:"-"`
    UpdatedAt   time.Time   `schema:"-"`
    CreatedAt   time.Time   `schema:"-"`
}

//Test represent an a test in a test list
type Test struct {
    ID         int         `schema:"-"`
    Expected   string      `schema:"exp"`
}

//Calculator represent the questions for the calculator
type Calculator struct {
    ID         int         `schema:"-"`
    Type       string      `schema:"type"`
    Amount     int         `schema: "amount"`
    Question   string      `schema:"quest"`
    Option01    string      `schema:"-"`
    Points01    int         `schema:"-"`
    Option02    string      `schema:"-"`
    Points02    int         `schema:"-"`
    Option03    string      `schema:"-"`
    Points03    int         `schema:"-"`
    Option04    string      `schema:"-"`
    Points04    int         `schema:"-"`
    Option05    string      `schema:"-"`
    Points05    int         `schema:"-"`
    Option06    string      `schema:"-"`
    Points06    int         `schema:"-"`
    Option07    string      `schema:"-"`
    Points07    int         `schema:"-"`
    Option08    string      `schema:"-"`
    Points08    int         `schema:"-"`
    Option09    string      `schema:"-"`
    Points09    int         `schema:"-"`
    Option10    string      `schema:"-"`
    Points10    int         `schema:"-"`
}
//Result represent the results from the calculator
type Result struct {
    ID         int         `schema:"-"`
    Total      int         `schema:"total"`
    TakenAt   time.Time   `schema:"-"`
    P01   string      `schema:"-"`
    P02   string      `schema:"-"`
    P03   string      `schema:"-"`
    P04   string      `schema:"-"`
    P05   string      `schema:"-"`
    P06   string      `schema:"-"`
    P07   string      `schema:"-"`
    P08   string      `schema:"-"`
    P09   string      `schema:"-"`
    P10   string      `schema:"-"`
    P11   string      `schema:"-"`
}


//TableName is a function to modify the standard name
//for the database
//func (*Results) TableName() {
    //return "calculator_res"
//}

