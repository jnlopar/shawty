package web

import (
	"go.3fps.com/utils/log"
	"html/template"
	"net/http"
)

var indexHtml *template.Template

// getIndexHtml returns the template for index.html
func getIndexHtml() *template.Template {
	if indexHtml == nil {
		indexHtml = template.Must(template.ParseFiles("templates/index.html"))
	}
	return indexHtml
}

// Structure for the index page
type HomeController struct {
	config map[string]string
}

// NewHomeController creates a new HomeController instance
func NewHomeController(config map[string]string) *HomeController {
	index := new(HomeController)
	index.config = config
	return index
}

// Respond creates the response package for the index page
func (page *HomeController) Index() *ResPkg {
	res := NewResPkg()
	res.Data = map[string]interface{}{"Domain": page.config["SHAWTY_DOMAIN"]}
	return res
}

func (page *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info("Index requested")
	defer closeReqBody(r)
	
	res := page.Index()
	tpl := getIndexHtml()
	if err := tpl.Execute(w, res); err != nil {
		log.Error("Cannot execute index template")
		log.Error(err)
		http.Error(w, "Error", http.StatusInternalServerError)		
	}
	
}
