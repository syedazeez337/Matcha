package index

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func LoadDocumentsFromDir(dir string) ([]Document, error) {
	var documents []Document

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	id := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !(strings.HasSuffix(name, ".txt") || strings.HasSuffix(name, ".md")) {
			continue
		}

		path := filepath.Join(dir, name)

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return nil, errors.New("failed to read " + path + ": " + err.Error())
		}

		doc := Document{
			ID: id,
			Title: name,
			Path: path,
			Content: string(contentBytes),
		}

		documents = append(documents, doc)
		id++
	}

	return documents, nil
}
