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

	fmt.Printf("Loaded %d documents:\n", len(docs))
	for _, doc := range docs {
		fmt.Printf("- [%d] %s (%s)\n", doc.ID, doc.Title, doc.Path)
	}
}
