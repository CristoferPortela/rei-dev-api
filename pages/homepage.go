package pages

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type route struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type sliderItem struct {
	Title            string `json:"title"`
	Image            string `json:"image"`
	Layout           int    `json:"layout"`
	Description      string `json:"description"`
	CallToAction     string `json:"call_to_action"`
	ImageDescription string `json:"alt"`
	ImageAuthor      string `json:"author"`
}

type Page struct {
	Title  string       `json:"title"`
	Routes []route      `json:"routes"`
	Slider []sliderItem `json:"slider"`
}

func HomePage(c echo.Context) error {

	var r []route
	var s []sliderItem
	r = append(r, route{Title: "Contato", URL: "contato"}, route{Title: "Ofertas", URL: "ofertas"})
	s = append(s, sliderItem{Title: "Bem vindo", Image: "src://oi", Layout: 1, Description: "<p>Bluuu</p>", CallToAction: "/home", ImageDescription: "alter", ImageAuthor: "In pexels"})
	p := &Page{
		Title:  "Home",
		Routes: r,
		Slider: s,
	}
	return c.JSON(http.StatusOK, p)
}
