package models

import (
	"io/ioutil"
	"strings"

	"github.com/blevesearch/bleve"
)

//STRUCT
type Card struct {
	Id      string   `json:"id"`
	Breads  []string `json:"breads"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
}

func InitCard(title string) Card {
	cardBts, _ := ioutil.ReadFile("api/cards/" + title)
	cardStr := string(cardBts)

	tokens := strings.Split(cardStr, "\n")

	content := strings.Join(tokens[2:], "\n")

	return Card{
		Id:      title,
		Breads:  strings.Split(tokens[0], ","),
		Title:   tokens[1],
		Content: content,
	}
}

// METODO UTILIZZATO PER AGGIUNGERE UNA CARTA ALL'INDEX
func (e *Card) Index(index bleve.Index) error {
	err := index.Index(string(e.Id), e)
	return err
}

//card.Index(idx)
