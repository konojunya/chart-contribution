package main

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Contribute struct {
	Date  string
	Count string
}

func getContribute(id string) []Contribute {
	doc, err := goquery.NewDocument("https://github.com/users/" + id + "/contributions")
	if err != nil {
		log.Fatal(err)
	}
	var contributions []Contribute

	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		date, _ := s.Attr("data-date")
		count, _ := s.Attr("data-count")
		contributions = append(contributions, Contribute{
			Date:  date,
			Count: count,
		})
	})

	return contributions
}

func main() {

	r := gin.Default()

	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")

	r.LoadHTMLGlob("view/*")

	// web
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// api
	r.GET("/api/contributions", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(http.StatusOK, getContribute(id))
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":" + os.Getenv("PORT"))

}
