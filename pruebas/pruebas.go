package main

type MarkdownFile struct {
	Title   string
	RelPath string
}

type Node struct {
	Name     string        // Nombre ("docs", "guia.md", etc.)
	IsDir    bool          // Si es directorio o no
	File     *MarkdownFile // {Title, RelPath} =>nil si es carpeta; si es archivo, apunta a sus datos
	Children []*Node       // Lista de punteros a sus hijos (carpetas o archivos)
}

// Busca si `n' ya tiene un hijo con ese nombre.
// Si existe, lo devuelve. Sino, lo crea, lo guarda en su hijos y lo devuelve
func (n *Node) getOrCreateChild(name string, isDir bool) *Node {
	// 1. Buscamos si ya existe entre los hijos actuales
	for _, child := range n.Children {
		if child.Name == name {
			return child // Lo encontramos, devolvemos el puntero o ese hijo
		}
	}

	// 2. Si llegamosa aqui, no existía. Creamos un nuevo nodo en memoria con &
	nuevo := &Node{
		Name:  name,
		IsDir: isDir,
	}

	// 3. Lo enganchamos a la lista de hijos del nodo actual
	n.Children = append(n.Children, nuevo)

	// 4. Devolvemos el puntero al nuevo nodo para poder avanzar hacia el
	return nuevo
}

func main() {
	var archivos = []MarkdownFile{
		{Title: "Página Principal", RelPath: "docs/markdown/index.md"},
		{Title: "Guía de Instalación", RelPath: "docs/markdown/guia-instalacion.md"},
		{Title: "Configuración del Sistema", RelPath: "config/settings.json"},
		{Title: "Variables de Entorno", RelPath: "config/.env.example"},
		{Title: "Esquema de Base de Datos", RelPath: "database/schema.sql"},
		{Title: "Documentación de la API", RelPath: "docs/markdown/api.md"},
		{Title: "Registro de Cambios", RelPath: "CHANGELOG.md"},
		{Title: "Licencia", RelPath: "LICENSE"},
	}
}
