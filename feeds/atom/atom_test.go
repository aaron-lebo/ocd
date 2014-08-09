package atom

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFeed(t *testing.T) {
	resp, err := http.Get("http://github.com/aaron-lebo.atom")
	if err != nil {
		t.Error("Error opening url:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		t.Error("Error parsing response:", err)
	} else if feed.Links[0].Href != "https://github.com/aaron-lebo" {
		t.Error("bad Feed.Links[0].Href")
	}
}
