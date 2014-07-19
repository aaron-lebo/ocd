package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/fisher-lebo/ocd/atom"
	"github.com/fisher-lebo/ocd/rss"
	"github.com/fisher-lebo/ocd/rss1"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error opening url:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var rss_ rss.Rss
	decoder := xml.NewDecoder(bytes.NewBuffer(body))
	decoder.DefaultSpace = "rss"
	err = decoder.Decode(&rss_)
	if err != nil {
		fmt.Println("Error parsing response:", err)
	} else {
		fmt.Println(rss_.Channel.Link)
		fmt.Printf("%+v\n", rss_.Channel.AtomLink)
		for _, item := range rss_.Channel.Items {
			fmt.Println(item.Title)
		}
		return
	}

	var feed atom.Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Error parsing response:", err)
	} else {
		for _, entry := range feed.Entries {
			fmt.Println(entry.Title)
		}
		return
	}

	var rdf rss1.RDF
	err = xml.Unmarshal(body, &rdf)
	if err != nil {
		fmt.Println("Error parsing response:", err)
	} else {
		for _, item := range rdf.Items {
			fmt.Println(item)
		}
	}
}
