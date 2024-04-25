package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func parseFiles(dirName string) map[string][]Link {
	files, _ := os.ReadDir(dirName)
	linksMap := make(map[string][]Link)

	for _, fileEntry := range files {
		file, _ := os.Open(dirName + "/" + fileEntry.Name())
		tokenizer := html.NewTokenizer(file)
		links := []Link{}

		for {
			tokenType := tokenizer.Next()
			if tokenType == html.ErrorToken {
				break
			}
			token := tokenizer.Token()
			if token.Data == "a" {
				tokenType := tokenizer.Next()
				if tokenType == html.TextToken {
					text := tokenizer.Token()
					href := ""
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							href = attr.Val
							break
						}
					}
					link := Link{
						Href: href,
						Text: text.String(),
					}
					links = append(links, link)
				}
			}
		}

		linksMap[fileEntry.Name()] = links
	}

	return linksMap
}

func main() {

	inputsDirName := "inputs"

	linksMap := parseFiles(inputsDirName)

	outPutFileName := "output.json"
	file , _ := os.Create(outPutFileName);
	defer file.Close()
	jsonData , _ := json.Marshal(linksMap);
	file.Write(jsonData);
	// Use the linksMap as needed
	fmt.Println(linksMap)

}
