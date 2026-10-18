# md-indexer

md-indexer-go/
├── main.go       // Solo orquesta el flujo general y valida los argumentos
├── markdown.go   // Todo lo relativo a archivos .md (struct MarkdownFile, extractTitle, walk)
└── tree.go       // Toda la estructura de datos del árbol (struct Node, GetOrCreateChild, etc.)
