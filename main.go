package main

import (
	"fmt"
	"os"
)

// Verificación de argumentos
func checkArgs() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Falta el nombre")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Warning: Se ha introducido más de 1 parámetro y se usará sólo el primero.")
	}
}

// Verificación del path
func checkPath() string {
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
	return sourcePath
}

func main() {
	checkArgs()
	sourcePath := checkPath()

	rutas, err := walkMarkdownFiles(sourcePath)
	if err != nil {
		fmt.Printf("Error al recorrer %s: %s\n", sourcePath, err)
		return
	}

	nodos := WalkPaths(rutas)
	textoFinal := Iterator(nodos) // Genero el texto con todo el árbol de ficheros
	fmt.Println(textoFinal)
	nombreFichero := "index.md"
	// StringToFile(targetDir string, file string, text string)
	StringToFile(sourcePath, nombreFichero, textoFinal) // Genero el fichero final
}
