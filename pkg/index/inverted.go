package index

import "github.com/syedazeez337/matcha/pkg/tokenize"

type InvertedIndex struct {
	Index map[string]map[int]struct{}
}

func NewInvertedIndex(docs []Document) *InvertedIndex {
	idx := &InvertedIndex{
		Index: make(map[string]map[int]struct{}),
	}

	for _, doc := range docs {
		tokens := tokenize.Tokenize(doc.Content)

		for _, token := range tokens {
			// Initialize set if needed
			if idx.Index[token] == nil {
				idx.Index[token] = make(map[int]struct{})
			}
			// Add doc ID to the set
			idx.Index[token][doc.ID] = struct{}{}
		}
	}

	return idx
}


func (idx *InvertedIndex) Search(token string) []int {
	token = tokenize.Tokenize(token)[0]

	docSet, found := idx.Index[token]
	if !found {
		return nil
	}

	result := make([]int, 0, len(docSet))
	for docID := range docSet {
		result = append(result, docID)
	}

	return result
}