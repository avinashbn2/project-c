package main

import (
	"bytes"
	"cproject/pkg/reader"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
)

type Obj struct {
	header string
	items  []string
	links  []string
}
type Payload struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Tag      string `json:"tag"`
	Author   string `json:"author"`
	Length   int    `json:"tag"`
	Excerpt  string `json:"excerpt"`
	SiteName string `json:"sitename"`
	Image    string `json:"image"`
}

const URL string = "https://github.com/lauragift21/awesome-learning-resources"

func main() {
	c := colly.NewCollector()
	objs := make([]Obj, 10)
	c.OnHTML("article ", func(e *colly.HTMLElement) {
		ulIndex := 0
		e.ForEach("article>  h2", func(i int, li *colly.HTMLElement) {
			if ulIndex == 0 {
				ulIndex = i + 1
			}
			header := e.ChildText(fmt.Sprintf(" article > h2:nth-of-type(%d)", i+1))
			items := e.ChildTexts(fmt.Sprintf("article > ul:nth-of-type(%d) a", ulIndex))
			links := e.ChildAttrs(fmt.Sprintf("article > ul:nth-of-type(%d) a", ulIndex), "href")
			items2 := e.ChildTexts(fmt.Sprintf("article > h2:nth-of-type(%d) +p > a", i+1))

			if len(items2) > 0 {
				objs = append(objs, Obj{header: header, items: items2, links: links})
			} else {
				objs = append(objs, Obj{header: header, items: items, links: links})
				ulIndex++
			}
		})
		for i, obj := range objs {
			for j, item := range obj.items {
				if i == 10 {
					continue
				}

				fetchItem, err := reader.Fetch(obj.links[j])
				if err != nil {
					fmt.Println("ERROR HERE", err)
					continue
				}

				if fetchItem == nil {
					continue
				}

				p := &Payload{Name: item, URL: obj.links[j], Tag: obj.header, Excerpt: fetchItem.Excerpt, Author: fetchItem.Author, Length: fetchItem.Length, Image: fetchItem.Image, SiteName: fetchItem.SiteName}
				data, err := json.Marshal(p)
				if err != nil {
					fmt.Println(err)
				}
				resp, err := http.Post("http://localhost:3001/resource", "applcation/json", bytes.NewBuffer(data))

				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("resp", resp)
				fmt.Println(i, j, item, obj.header, obj.links[j])
			}
		}
	})
	c.Visit(URL)

}
