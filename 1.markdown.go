package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type MarkdownFile struct {
	Title   string
	RelPath string
}

// walkMarkdownFiles recorre sourcePath buscando ficheros .md y devuelve sus
// datos (título + ruta relativa). Ignora los index.md y los directorios.
func walkMarkdownFiles(sourcePath string) ([]*MarkdownFile, error) {
	var mdFiles []*MarkdownFile
	err := filepath.WalkDir(sourcePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".md" {
			return nil
		}
		cleanPath := filepath.ToSlash(path)
		// Obtengo la ruta relativa
		relPath, err := filepath.Rel(sourcePath, cleanPath)
		if err != nil {
			return err
		}
		cleanRelPath := filepath.ToSlash(relPath)
		if strings.ToLower(cleanRelPath) == "index.md" {
			return nil
		}
		title, err := extractTitle(cleanPath)
		if err != nil {
			fmt.Println("Ha habido algun error")
		}
		mdFiles = append(mdFiles, &MarkdownFile{Title: title, RelPath: cleanRelPath})
		return nil // Continuar con el siguiente elemento
	})
	return mdFiles, err
}

func extractTitle(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if rest, found := strings.CutPrefix(line, "# "); found {
			return rest, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	// Si no se ha devuelto el título
	namefile := strings.TrimSuffix(filepath.Base(path), ".md")
	return namefile, nil
}
