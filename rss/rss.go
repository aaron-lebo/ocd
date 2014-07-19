package rss

import (
	"encoding/xml"
	"github.com/fisher-lebo/ocd/atom"
	"time"
)

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

type Hour struct {
	Text string `xml:",chardata"`
}

type SkipHours struct {
	Hours []Hour `xml:"hour"`
}

type Day struct {
	Text string `xml:",chardata"`
}

type SkipDays struct {
	Days []Day `xml:"day"`
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

type Item struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Author      string     `xml:"author"`
	Categories  []Category `xml:"category"`
	Comments    string     `xml:"comments"`
	Enclosure   Enclosure  `xml:"enclosure"`
	Guid        Guid       `xml:"guid"`
	PubDate     string     `xml:"pubDate"`
	Source      Source     `xml:"source"`
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
	skipHours     SkipHours  `xml:"skipHours"`
	skipDays      SkipDays   `xml:"skipDays"`
	Items         []Item     `xml:"item"`
	AtomLink      atom.Link  `xml:"http://www.w3.org/2005/Atom link"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}
