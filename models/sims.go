package models

import (
	"time"
)

//Person represent an item of search list
type Search struct {
	ID        int       `schema:"-"`
	Title     string    `schema:"title"`
	Body      string    `schema:"body"`
	URL       string    `schema:"url"`
	Dom       string    `schema:"dom"`
	DomTitle  string    `schema:"-"`
	UpdatedAt time.Time `schema:"-"`
	CreatedAt time.Time `schema:"-"`
}

//Test represent an a test in a test list
type Test struct {
	ID       int    `schema:"-"`
	Expected string `schema:"exp"`
}

//Calculator represent the questions for the calculator
type Calculator struct {
<<<<<<< HEAD
	ID       int     `schema:"-"`
	Type     string  `schema:"type"`
	Amount   int     `schema: "amount"`
	Question string  `schema:"quest"`
	Recomm   string  `schema:"-"`
	Recomm1  string  `schema:"-"`
	Recomm2  string  `schema:"-"`
	Recomm3  string  `schema:"-"`
	Recomm4  string  `schema:"-"`
	Recomm5  string  `schema:"-"`
	Option01 string  `schema:"-"`
	Points01 float32 `schema:"-"`
	Option02 string  `schema:"-"`
	Points02 float32 `schema:"-"`
	Option03 string  `schema:"-"`
	Points03 float32 `schema:"-"`
	Option04 string  `schema:"-"`
	Points04 float32 `schema:"-"`
	Option05 string  `schema:"-"`
	Points05 float32 `schema:"-"`
	Option06 string  `schema:"-"`
	Points06 float32 `schema:"-"`
	Option07 string  `schema:"-"`
	Points07 float32 `schema:"-"`
	Option08 string  `schema:"-"`
	Points08 float32 `schema:"-"`
	Option09 string  `schema:"-"`
	Points09 float32 `schema:"-"`
	Option10 string  `schema:"-"`
	Points10 float32 `schema:"-"`
=======
    ID         int         `schema:"-"`
    Type       string      `schema:"type"`
    Amount     int         `schema: "amount"`
    Question   string      `schema:"quest"`
    Recomm     string     `schema:"-"`
    Recomm1     string     `schema:"-"`
    Recomm2     string     `schema:"-"`
    Recomm3     string     `schema:"-"`
    Recomm4     string     `schema:"-"`
    Recomm5     string     `schema:"-"`
    Option01    string      `schema:"-"`
    Points01    float32     `schema:"-"`
    Option02    string      `schema:"-"`
    Points02    float32     `schema:"-"`
    Option03    string      `schema:"-"`
    Points03    float32     `schema:"-"`
    Option04    string      `schema:"-"`
    Points04    float32     `schema:"-"`
    Option05    string      `schema:"-"`
    Points05    float32     `schema:"-"`
    Option06    string      `schema:"-"`
    Points06    float32     `schema:"-"`
    Option07    string      `schema:"-"`
    Points07    float32     `schema:"-"`
    Option08    string      `schema:"-"`
    Points08    float32     `schema:"-"`
    Option09    string      `schema:"-"`
    Points09    float32     `schema:"-"`
    Option10    string      `schema:"-"`
    Points10    float32     `schema:"-"`
>>>>>>> dfd75cfa71185a4f227dc93d9c9352a304562b4a
}

//Result represent the results from the calculator
type Result struct {
	ID        int       `schema:"-"`
	Total     float32   `schema:"total"`
	CreatedAt time.Time `schema:"-"`
	Recomm    string    `schema:"-"`
}

//FunFact represent the results from the calculator
type FunFact struct {
	ID       int    `schema:"-"`
	Category string `schema:"-"`
	Fact     string `schema:"-"`
	URL      string `schema:"-"`
}

//TableName is a function to modify the standard name
//for the database
//func (*Results) TableName() {
//return "calculator_res"
//}
