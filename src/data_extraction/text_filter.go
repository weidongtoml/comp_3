// Copyright Weidong Liang 2013. All rights reserved.

package main

import (
	"code.google.com/p/go.net/html"
	"log"
	"strings"
)

func StripCodeSectionFromHTML(html_text string) string {
	doc, err := html.Parse(strings.NewReader("<html>" + html_text + "</html>"))
	if err != nil {
		log.Printf("Failed to parse html content: %s", html_text)
		return ""
	}
	return stripCodeSection(doc)
}

func stripCodeSection(node *html.Node) string {
	result_str := ""
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if c.Data == "pre" || c.Data == "code" {
				//Skip pre and code sections
			} else {
				result_str += stripCodeSection(c) + " "
			}
		} else if c.Type == html.TextNode {
			result_str += strings.Map(func(r rune) rune {
				var c rune
				switch r {
				case '\r', '\t', '\n', '\f':
					c = ' '
					break
				default:
					c = r
					break
				}
				return c
			}, c.Data) + " "
		} else {
			// Skip the rest?
		}
	}
	return result_str
}
