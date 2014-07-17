package rss1

import (
	"encoding/xml"
	"time"
)

type DublinCore struct {
	Title       string    `xml:"http://purl.org/dc/elements/1.1/ title"`
	Creator     string    `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Subject     string    `xml:"http://purl.org/dc/elements/1.1/ subject"`
	Description string    `xml:"http://purl.org/dc/elements/1.1/ description"`
	Contributor string    `xml:"http://purl.org/dc/elements/1.1/ contributor"`
	Date        time.Time `xml:"http://purl.org/dc/elements/1.1/ date"`
	Type        string    `xml:"http://purl.org/dc/elements/1.1/ type"`
	Format      string    `xml:"http://purl.org/dc/elements/1.1/ format"`
	Identifier  string    `xml:"http://purl.org/dc/elements/1.1/ identifier"`
	Source      string    `xml:"http://purl.org/dc/elements/1.1/ source"`
	Language    string    `xml:"http://purl.org/dc/elements/1.1/ language"`
	Relation    string    `xml:"http://purl.org/dc/elements/1.1/ relation"`
	Coverage    string    `xml:"http://purl.org/dc/elements/1.1/ coverage"`
	Rights      string    `xml:"http://purl.org/dc/elements/1.1/ rights"`
}

type Syndication struct {
	UpdatePeriod    string    `xml:"http://purl.org/1.0/modules/syndication/ updatePeriod"`
	UpdateFrequency int       `xml:"http://purl.org/1.0/modules/syndication/ updateFrequency"`
	UpdateBase      time.Time `xml:"http://purl.org/1.0/modules/syndication/ updateBase"`
}

type Image struct {
	DublinCore
	About string `xml:"about,attr"`
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Link  string `xml:"link"`
}

type Li struct {
	Resource string `xml:"resource,attr"`
}

type TextInput struct {
	DublinCore
	About       string `xml:"about,attr"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

type Item struct {
	DublinCore
	About       string `xml:"about,attr"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type Channel struct {
	DublinCore
	Syndication
	About       string    `xml:"about,attr"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Image       Image     `xml:"image"`
	Items       []Li      `xml:"items>Seq>li"`
	TextInput   TextInput `xml:"textinput"`
}

type RDF struct {
	XMLName xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	Channel Channel  `xml:"channel"`
	Items   []Item   `xml:"item"`
}
