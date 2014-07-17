package rss1

import (
	"encoding/xml"
	"time"
)

type DublinCore struct {
	Title       string    `xml:"title"`
    Creator     string    `xml:"creator"`
	Subject     string    `xml:"subject"`
	Description string    `xml:"description"`
	Contributor string    `xml:"contributor"`
	Date        time.Time `xml:"date"`
	Type        string    `xml:"type"`
	Format      string    `xml:"format"`
	Identifier  string    `xml:"identifier"`
	Source      string    `xml:"source"`
	Language    string    `xml:"language"`
	Relation    string    `xml:"relation"`
	Coverage    string    `xml:"coverage"`
	Rights      string    `xml:"rights"`
}

type Syndication struct {
	UpdatePeriod    string    `xml:"updatePeriod"`
	UpdateFrequency int       `xml:"updateFrequency"`
	UpdateBase      time.Time `xml:"updateBase"`
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
