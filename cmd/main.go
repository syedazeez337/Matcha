package main

import (
	"fmt"

	"github.com/syedazeez337/matcha/pkg/index"
	"github.com/syedazeez337/matcha/pkg/tokenize"
)

func main() {
	docs, err := index.LoadDocumentsFromDir("docs")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %d documents:\n", len(docs))
	for _, doc := range docs {
		fmt.Printf("- [%d] %s (%s)\n", doc.ID, doc.Title, doc.Path)

		tokens := tokenize.Tokenize(doc.Content)
		fmt.Printf("  First 10 tokens: %v\n", tokens[:10])
	}
}
