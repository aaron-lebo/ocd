package rss

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRss(t *testing.T) {
	resp, err := http.Get("http://reddit.com/.rss")
	if err != nil {
		t.Error("Error opening url:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var rss Rss
	decoder := xml.NewDecoder(bytes.NewBuffer(body))
	decoder.DefaultSpace = "rss"
	err = decoder.Decode(&rss)
	if err != nil {
		t.Error("Error parsing response:", err)
	} else if rss.Channel.Link != "http://www.reddit.com/" {
		t.Error("bad Rss.Channel.Link")
	}
}
