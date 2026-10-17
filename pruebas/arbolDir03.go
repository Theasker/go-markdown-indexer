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

type Node struct {
	Name     string
	IsDir    bool
	File     *MarkdownFile
	Children []*Node
}

func (n *Node) GetOrCreateChild(name string) *Node {
	// 1. ¿Ya existe este hijo en la lista?
	for _, hijo := range n.Children {
		if hijo.Name == name {
			return hijo
		}
	}
	// 2. Si no existe creamos uno nuevo con &
	nuevo := &Node{
		Name:     name,
		IsDir:    true,
		File:     nil,
		Children: nil,
	}
	// 3. Lo agregamos como hijo de la carpeta actual
	n.Children = append(n.Children, nuevo)
	// 4. Devolvemos el puntero al nuevo hijo
	return nuevo
}

func main() {
	var rutas = []MarkdownFile{
		{Title: "Página Principal", RelPath: "docs/markdown/index.md"},
		{Title: "Guía de Instalación", RelPath: "docs/markdown/guia-instalacion.md"},
		{Title: "Configuración del Sistema", RelPath: "config/settings.json"},
		{Title: "Variables de Entorno", RelPath: "config/.env.example"},
		{Title: "Esquema de Base de Datos", RelPath: "database/schema.sql"},
		{Title: "Documentación de la API", RelPath: "docs/markdown/api.md"},
		{Title: "Registro de Cambios", RelPath: "CHANGELOG.md"},
		{Title: "Licencia", RelPath: "LICENSE"},
	}

	raiz := &Node{Name: "root", IsDir: true}
	cursor := raiz

	for _, ruta := range rutas {
		partesRuta := strings.Split(ruta.RelPath, "/")
		cursor = raiz // Al terminar cada iteración, volvemos a la raíz
		for idx, parte := range partesRuta {
			// Si es la última parte de la ruta = fichero
			fmt.Println(partesRuta, "->", parte, "->", len(partesRuta)-1, idx)
			copiaRuta := &ruta // Hacemos una copia local única para este archivo
			if (len(partesRuta) - 1) == idx {
				newFile := &Node{
					Name:     parte,
					IsDir:    false,
					File:     copiaRuta,
					Children: nil,
				}
				cursor.Children = append(cursor.Children, newFile)
			} else { // Si no es el fichero final de la ruta
				cursor = cursor.GetOrCreateChild(parte)
			}
		}
	}
	//odump.Dump("-------------------------", raiz)
	godump.DumpJSON(raiz)
}
