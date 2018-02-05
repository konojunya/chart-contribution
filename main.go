package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Contribute struct {
	Date  string
	Count string
}

var contributions []Contribute

func init() {
	doc, err := goquery.NewDocument("https://github.com/users/konojunya/contributions")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("rect").Each(func(_ int, s *goquery.Selection) {
		date, _ := s.Attr("data-date")
		count, _ := s.Attr("data-count")
		contributions = append(contributions, Contribute{
			Date:  date,
			Count: count,
		})
	})
}

func main() {

	r := gin.Default()
	r.GET("/api/contributions", func(c *gin.Context) {
		c.JSON(http.StatusOK, contributions)
	})

	r.Run(":8000")

}
