package rss1

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRss(t *testing.T) {
	resp, err := http.Get("https://feeds.pinboard.in/rss/u:fisher.lebo/")
	if err != nil {
		t.Error("Error opening url:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var rdf RDF
	err = xml.Unmarshal(body, &rdf)
	if err != nil {
		t.Error("Error parsing response:", err)
	} else if rdf.Channel.Link != "https://pinboard.in/u:fisher.lebo/public/" {
		t.Error("bad RDF.Channel.Link")
	}
}
