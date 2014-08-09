package rss

import (
	"encoding/xml"
	"github.com/aaron-lebo/ocd/feeds/modules/atom"
	"github.com/aaron-lebo/ocd/feeds/modules/media"
	"time"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string     `xml:"title"`
	Link          string     `xml:"rss link"`
	Description   string     `xml:"description"`
	Language      string     `xml:"language"`
	Copyright     string     `xml:"copyright"`
	WebMaster     string     `xml:"webMaster"`
	pubDate       time.Time  `xml:"pubDate"`
	LastBuildDate time.Time  `xml:"lastBuildDate"`
	Categories    []Category `xml:"category"`
	Generator     string     `xml:"generator"`
	Docs          string     `xml:"docs"`
	Cloud         Cloud      `xml:"cloud"`
	Ttl           string     `xml:"ttl"`
	Image         Image      `xml:"image"`
	Rating        string     `xml:"rating"`
	TextInput     TextInput  `xml:"textInput"`
	SkipHours     []int      `xml:"skipHours>hour"`
	SkipDays      []string   `xml:"skipDays>day"`
	Items         []Item     `xml:"item"`
	atom.Atom
}

type Category struct {
	Domain string `xml:"domain,attr"`
	Text   string `xml:",chardata"`
}

type Cloud struct {
	Domain            string `xml:"domain,attr"`
	Port              string `xml:"port,attr"`
	Path              string `xml:"path,attr"`
	RegisterProcedure string `xml:"registerProcedure,attr"`
	Protocol          string `xml:"protocol,attr"`
}

type Image struct {
	Url         string `xml:"url"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Width       string `xml:"width"`
	Height      string `xml:"height"`
	Description string `xml:"description"`
}

type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

type Item struct {
	Title       string     `xml:"rss title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Author      string     `xml:"author"`
	Categories  []Category `xml:"category"`
	Comments    string     `xml:"comments"`
	Enclosure   Enclosure  `xml:"enclosure"`
	Guid        Guid       `xml:"guid"`
	PubDate     string     `xml:"pubDate"`
	Source      Source     `xml:"source"`
	media.Media
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	Text   string `xml:",chardata"`
}

type Guid struct {
	PermaLink string `xml:"permaLink,attr"`
	Text      string `xml:",chardata"`
}

type Source struct {
	Url  string `xml:"url,attr"`
	Text string `xml:",chardata"`
}
