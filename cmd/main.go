package main

import (
	"fmt"

	"github.com/syedazeez337/matcha/pkg/index"
)

func main() {
	docs, err := index.LoadDocumentsFromDir("docs")
	if err != nil {
		panic(err)
	}

	/*
		fmt.Printf("Loaded %d documents:\n", len(docs))
		for _, doc := range docs {
			fmt.Printf("- [%d] %s (%s)\n", doc.ID, doc.Title, doc.Path)

			tokens := tokenize.Tokenize(doc.Content)
			fmt.Printf("  First 10 tokens: %v\n", tokens[:10])
		}
	*/
	fmt.Printf("Loaded %d documents.\n", len(docs))

	//build the index
	invIdx := index.NewInvertedIndex(docs)

	// test query
	query := "language"
	matched := invIdx.Search(query)

	fmt.Printf("Documents containing %s:\n", query)
	for _, id := range matched {
		fmt.Printf("- [%d] %s\n", docs[id].ID, docs[id].Title)
	}
}
