package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func readHTML(filePath string) (*html.Node, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func stripAttributes(n *html.Node) {
	if n.Type == html.ElementNode {
		var newAttrs []html.Attribute
		for _, attr := range n.Attr {
			if strings.HasPrefix(attr.Key, "data-") {
				newAttrs = append(newAttrs, attr)
			}
		}
		n.Attr = newAttrs
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stripAttributes(c)
	}
}

func renderHTML(n *html.Node, writer *os.File) error {
	err := html.Render(writer, n)
	if err != nil {
		return fmt.Errorf("error rendering HTML: %v", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <inputfile> <outputfile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	doc, err := readHTML(inputFile)
	if err != nil {
		fmt.Printf("Error reading HTML: %v\n", err)
		return
	}

	stripAttributes(doc)

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	err = renderHTML(doc, outFile)
	if err != nil {
		fmt.Printf("Error rendering HTML: %v\n", err)
		return
	}

	fmt.Println("HTML processing complete. Output saved to", outputFile)
}
