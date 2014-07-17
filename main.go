package main

import (
	"encoding/xml"
	"fmt"
	"github.com/fisher-lebo/ocd/atom"
	"github.com/fisher-lebo/ocd/rss"
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
	err = xml.Unmarshal(body, &rss_)
	if err != nil {
		fmt.Println("Error parsing response:", err)
	} else {
		for _, item := range rss_.Channel.Items {
			fmt.Println(item.Title)
		}
		return
	}

	var feed atom.Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}
    for _, entry := range feed.Entries {
        fmt.Println(entry.Title)
    }
}
