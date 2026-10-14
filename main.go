package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

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

	// Recorrer el path
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
			println("[Markdown válido]: " + cleanPath)
		}

		return nil // Continuar con el siguiente elemento
	})

	if err != nil {
		fmt.Println("Error al recorrer:", err)
	}
}
