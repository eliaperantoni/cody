package models

import (
	"bufio"
	"os"

	"github.com/blevesearch/bleve"
)

//STRUCT
type Card struct {
	Id      string `json:"id"`
	Breads  string `json:"breads"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func InitCard(title string) Card {
	file, _ := os.Open("cards/" + title)
	fscanner := bufio.NewScanner(file)
	var i Card
	i.Id = title
	cont := 0
	for fscanner.Scan() {
		switch cont {
		case 0:
			//BREADS
			i.Breads = fscanner.Text()
		case 1:
			//TITLE
			i.Title = fscanner.Text()
		case 2:
			//CONTENT
			i.Content = fscanner.Text()
		default:
		}
		cont++
	}
	return i
}

// METODO UTILIZZATO PER AGGIUNGERE UNA CARTA ALL'INDEX
func (e *Card) Index(index bleve.Index) error {
	err := index.Index(string(e.Id), e)
	return err
}

//card.Index(idx)
