package conn

import (
	"os"

	"github.com/blevesearch/bleve"
)

var bleveIdx bleve.Index

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

// FUNZIONE CHE SI CONNETTE AD UN INDEX O CREA UN'ISTANZA
func Bleve(indexPath string) (bleve.Index, error) {

	// bleveIdx NON INIZIALIZZATO
	if bleveIdx == nil {
		var err error
		// PROVO AD APRIRLO
		bleveIdx, err = bleve.Open(indexPath)
		if err != nil {
			//CREO NUOVO INDEX
			mapping := bleve.NewIndexMapping()
			bleveIdx, err = bleve.New(indexPath, mapping)
			if err != nil {
				return nil, err
			}
		}
	}

	return bleveIdx, nil
}

func BlaveIndex(testIdx string) (bleve.Index, error) {
	idx, err := Bleve(testIdx)
	return idx, err
}

func FindByText(idx bleve.Index, text string) (*bleve.SearchResult , error) {
	
	query := bleve.NewMatchQuery(text)
	searchRequest := bleve.NewSearchRequest(query)
	search, err := idx.Search(searchRequest)

	return search, err
}
func idxDestroy(testIdx string) {
	os.RemoveAll(testIdx)
}
