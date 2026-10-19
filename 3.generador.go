package main

import (
	"fmt"
	"strings"
)

func renderizar(elem *Node, deepth int, sb *strings.Builder) {
	sangria := strings.Repeat("  ", deepth)

	if elem.IsDir {
		sb.WriteString(fmt.Sprintf("%s- **%s**\n", sangria, elem.Name))
	} else {
		sb.WriteString(fmt.Sprintf("%s- [%s](%s)\n", sangria, elem.File.Title, elem.File.RelPath))
	}

	for _, hijo := range elem.Children {
		renderizar(hijo, deepth+1, sb)
	}
}

func Iterator(nodos *Node) string {
	var builder strings.Builder
	for _, nodo := range nodos.Children {
		renderizar(nodo, 0, &builder)
	}

	markdown := builder.String()
	return markdown
}
