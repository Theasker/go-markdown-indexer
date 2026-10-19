package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func StringToFile(targetDir string, file string, text string) {

	rutaDestino := filepath.Join(targetDir, file)
	err := os.WriteFile(rutaDestino, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo:", err)
		return
	}
	fmt.Print("Archivo '", file, "' creado con éxito en: '", rutaDestino, "'")
}
