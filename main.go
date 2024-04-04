package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	//	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func walkDirWithExtension(path string, ext string) ([]string, error) {
	var targetFiles []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ext {
			_, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			targetFiles = append(targetFiles, path)

		}
		return nil

	})
	return targetFiles, err
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)

}

func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func main() {
	mdDir := "./md"
	publicDir := "./public/"

	files, err := walkDirWithExtension(mdDir, ".md")
	if err != nil {
		fmt.Println("Error walking directory: ", err)
	}
	fmt.Println(files)
	for _, file := range files {
		fileData, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Unable to read file: ", err)
			continue
		}
		htmlData := mdToHTML(fileData)
		fmt.Println(string(htmlData))
		_, fileName := filepath.Split(file)
		fileName = fileNameWithoutExtTrimSuffix(fileName)

		os.WriteFile(publicDir+fileName+".html", htmlData, 0644)

	}

}
