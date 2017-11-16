package models

import (
    "time"
)

//Todo represent an item of todo list
type Todo struct {
    ID          int         `schema:"-"`
    Title       string      `schema:"title"`
    Body        string      `schema:"body"`
    URL	        string      `schema:"url"`
    Dom         string       `schema:"dom"`
    CreatedAt   time.Time   `schema:"-"`
    UpdatedAt   time.Time   `schema:"-"`
}
