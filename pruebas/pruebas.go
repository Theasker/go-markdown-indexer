package main

import (
	"fmt"
	"strings"

	"github.com/goforj/godump"
)

type MarkdownFile struct {
	Title   string
	RelPath string
}

type Elemento struct {
	Nombre string
	Hijos  []*Elemento
}

func main() {
	var archivos = []MarkdownFile{
		{Title: "Página Principal", RelPath: "docs/markdown/index.md"},
		{Title: "Guía de Instalación", RelPath: "docs/markdown/guia-instalacion.md"},
		{Title: "Configuración del Sistema", RelPath: "config/settings.json"},
		{Title: "Variables de Entorno", RelPath: "config/.env.example"},
		{Title: "Esquema de Base de Datos", RelPath: "database/schema.sql"},
		{Title: "Punto de Entrada", RelPath: "src/main.go"},
		{Title: "Utilidades de Red", RelPath: "src/utils/network.go"},
		{Title: "Controlador de Rutas", RelPath: "src/handlers/routes.go"},
		{Title: "Pruebas Unitarias", RelPath: "tests/main_test.go"},
		{Title: "Documentación de la API", RelPath: "docs/markdown/api.md"},
		{Title: "Registro de Cambios", RelPath: "CHANGELOG.md"},
		{Title: "Licencia", RelPath: "LICENSE"},
		{Title: "Guía de Contribución", RelPath: "CONTRIBUTING.md"},
	}
	raiz := &Elemento{Nombre: "raiz"} // Creamos el nodo raíz usando & para obtener su puntero

	ruta := "fotos/viaje"
	partes := strings.Split(ruta, "/") // ["fotos", "viaje"]

	// Agregamos el primer nivel
	nodoFotos := &Elemento{Nombre: partes[0]}
	raiz.Hijos = append(raiz.Hijos, nodoFotos)

	// Agregamos el segundo nivel dentro de fotos
	nodoViaje := &Elemento{Nombre: partes[1]}
	nodoFotos.Hijos = append(nodoFotos.Hijos, nodoViaje)

	fmt.Printf("Raíz: %s -> Hijo: %s -> Subhijo: %s\n",
		raiz.Nombre, raiz.Hijos[0].Nombre, raiz.Hijos[0].Hijos[0].Nombre)
	godump.Dump(raiz)
}
