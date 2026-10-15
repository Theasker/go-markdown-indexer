package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/goforj/godump"
)

type MarkdownFile struct {
	Title   string
	RelPath string
}

type Node struct {
	Name     string
	IsDir    bool
	File     *MarkdownFile
	Children []*Node
}

func main() {
	// Verificación de argumentos
	if len(os.Args) < 2 {
		fmt.Println("Error: Falta el nombre")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Warning: Se ha introducido más de 1 parámetro y se usará sólo el primero.")
	}

	// Verificación del path
	sourcePath := os.Args[1]
	fmt.Println("Ruta recibida para indexar:", sourcePath)
	info, err := os.Stat(sourcePath)
	if err != nil {
		fmt.Println("El path", sourcePath, "no es un directorio válido:", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Println("Error: La ruta especificada no es un directorio")
		os.Exit(1)
	}

	// Recorrer el árbol completo
	var mdFiles []MarkdownFile
	err = filepath.WalkDir(sourcePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			//println("<dir>", path)
			return nil
		} else {
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
			mdFiles = append(mdFiles, MarkdownFile{Title: title, RelPath: cleanRelPath})
		}
		return nil // Continuar con el siguiente elemento
	})
	// godump.Dump(mdFiles)
	recorrerFicheros(mdFiles)

	if err != nil {
		fmt.Println("Error al recorrer:", err)
	}
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

func (n *Node) getOrCreateChild(name string, isDir bool) *Node {
	panic("no codificado")
}

func addFileToTree(root *Node, file MarkdownFile) {
	var splitPath []string = strings.Split(file.RelPath, "/")
	for idx, part_of_route := range splitPath {
		if (len(part_of_route) - 1) == idx {
			panic("no implementado")
		}
	}
}

func recorrerFicheros(mdFiles []MarkdownFile) {
	for _, item := range mdFiles {
		var partPath []string = strings.Split(item.RelPath, "/")
		godump.Dump(item, partPath)
	}
}
