package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID            string       `json:"id"`
	Contributions []Contribute `json:"contributions"`
}

type Contribute struct {
	Date  string `json:"date"`
	Count string `json:"count"`
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
		usersStr := c.Query("users")
		if len(usersStr) == 0 {
			c.JSON(http.StatusOK, nil)
			return
		}

		userIds := strings.Split(usersStr, " ")
		res := make(chan User, len(userIds))
		users := make([]User, 0)

		for _, userID := range userIds {
			go func(id string, res chan User) {
				log.Printf("Now Loading userID: %v\n", id)
				res <- getContribute(id)
			}(userID, res)
		}
		for _ = range userIds {
			user := <-res
			users = append(users, user)
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	r.Run(":" + os.Getenv("PORT"))

}

func getContribute(id string) User {
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

	user := User{
		ID:            id,
		Contributions: contributions,
	}

	return user
}
