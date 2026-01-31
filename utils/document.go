package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// Searchable interface - implement this to make any type searchable
type Searchable interface {
	GetID() int
	GetSearchText() string
}

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// Implement Searchable interface for Document
func (d Document) GetID() int {
	return d.ID
}

func (d Document) GetSearchText() string {
	return d.Text
}

func LoadDocuments(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	dec := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
