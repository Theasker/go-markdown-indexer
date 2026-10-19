package main

import (
	"strings"
)

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

func WalkPaths(files []*MarkdownFile) *Node {
	raiz := &Node{Name: "Raiz", IsDir: true}
	cursor := raiz // Comenzamos con la raiz del árbol

	for _, ruta := range files {
		partesRuta := strings.Split(ruta.RelPath, "/")
		cursor = raiz // Al terminar cada iteración, volvemos a la raíz
		for idx, parte := range partesRuta {
			copiaRuta := ruta // Hacemos una copia local única para este archivo
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
	return raiz
}
