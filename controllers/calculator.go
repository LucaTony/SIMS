package controllers

import (
	"fmt"

	"github.com/LucaTony/SIMS/models"
)

type Quiz struct {
	Question string
	Option1  string
	Option2  string
	Option3  string
	Option4  string
}

//CalcPost function to get questions and options from db
// and then display it on the website.
func (S *Search) CalcPost() {
	fmt.Println("CalcPost")
	quiz := []*models.Calculator{}

	//container type struct, struct Quiz from top of file
	var quest = []Quiz{}

	var option = []Quiz{}

	//declaration and initilalization of variables
	question := Quiz{}
	options := Quiz{}

	for i := range quiz {
		//Get the question from the db.
		question = Quiz{
			Question: quiz[i].Question,
		}
		for j := range quiz {
			//Get the responding answer options to each question.
			options = Quiz{
				Option1: quiz[j].Option01,
				Option2: quiz[j].Option02,
				Option3: quiz[j].Option03,
				Option4: quiz[j].Option04,
			}
		}
	}
	quest = append(quest, question)
	option = append(option, options)
}
