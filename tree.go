package main

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

func addFileToTree(root *Node, file MarkdownFile) {
	panic("")
}
