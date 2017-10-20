package main

import (
	"fmt"
	"log"
	"net/http"

	c "github.com/LucaTony/SIMS/controllers"
	"github.com/LucaTony/SIMS/models"
	"github.com/gernest/utron"
)


func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Todo{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(c.NewTodo)

	//404 custom handler
	app.SetNotFoundHandler(http.HandlerFunc(c.NotFound))

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
