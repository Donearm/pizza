package main

////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2019-2020, Gianluca Fiore
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
////////////////////////////////////////////////////////////////////////////////

// ALPHA: incomplete parser of README.md file to html and then GeoJson


import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"golang.org/x/net/html"
	"gitlab.com/golang-commonmark/markdown"
)

// the input file
const doc = "README.md"

// slice of strings to contain the data from the markdow->html conversion
var arr []string

// open input files and return a slice of bytes
func readFromFile(fn string) ([]byte, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

// extract text from each Markdown tokens, recursively
func extractText(tok markdown.Token) string {
	switch tok := tok.(type) {
	case *markdown.Text:
		return tok.Content
	case *markdown.Inline:
		text := ""
		for _, tok := range tok.Children {
			text += extractText(tok)
		}
		return text
	}
	return ""
}

// parse each HTML node in the documetn
func parseNode(n *html.Node) {
	// catch links
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if a.Val != "#address" {
					arr = append(arr, a.Val)
				}
			}
		}
	// catch text
	} else if n.Type == html.TextNode && n.Data != "" {
		if len(n.Data) >= 2 {
			arr = append(arr, n.Data)
		}
	}
	// recursively parse siblings, if any
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseNode(c)
	}
}

func main() {
	var title, rendererOutput string

	// read file
	data, err := readFromFile(doc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// parse each Markdown token
	md := markdown.New()
	tokens := md.Parse(data)
	if len(tokens) > 0 {
		if heading, ok := tokens[0].(*markdown.HeadingOpen); ok {
			for i := 1; i < len(tokens); i++ {
				if tok, ok := tokens[i].(*markdown.HeadingClose); ok && tok.Lvl == heading.Lvl {
					break
				}
				title += extractText(tokens[i])
			}
			title = strings.TrimSpace(title)
		}
	}

	// get html from Markdown
	rendererOutput = md.RenderTokensToString(tokens)

	// parse the html
	doc, _ := html.Parse(strings.NewReader(rendererOutput))
	parseNode(doc)

	fmt.Println(arr)

}

