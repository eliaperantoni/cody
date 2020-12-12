package main

import (
	"api/conn"
	"api/models"
	"github.com/blevesearch/bleve"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var idx bleve.Index

func searchHandler(ctx echo.Context) error {
	query := ctx.QueryParam("q")

	res, err := conn.FindByText(idx, query)
	if err != nil {
		return err
	} else if res != nil {
		var cards []models.Card

		for _, hit := range res.Hits {
			var breads []string

			for _, bread := range hit.Fields["breads"].([]interface{}) {
				breads = append(breads, bread.(string))
			}

			cards = append(cards, models.Card{
				Id:      hit.Fields["id"].(string),
				Title:   hit.Fields["title"].(string),
				Content: hit.Fields["content"].(string),
				Breads:  breads,
			})
		}

		return ctx.JSON(http.StatusOK, cards)
	}

	return nil
}

func dieOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initCards() error {
	var err error

	idx, err = conn.BlaveIndex("api/cody.index")
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir("api/cards")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		d := models.InitCard(entry.Name())
		err = d.Index(idx)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := initCards()
	dieOnErr(err)

	e := echo.New()

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "web/dist",
		HTML5: true,
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/api")
		},
	}))

	api := e.Group("/api")
	api.GET("/search", searchHandler)

	logrus.Fatal(e.Start(":5000"))
}
