package handlers

import (
	"github.com/LucaTony/SIMS/libhttp"
	"html/template"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/video.html.tmpl", "templates/model.html.tmpl", "templates/calculator.html.tmpl", "templates/eu.html.tmpl", "templates/footer.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
