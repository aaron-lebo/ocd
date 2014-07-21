package rss1

import (
	"encoding/xml"
	"github.com/fisher-lebo/ocd/modules/dublincore"
	"github.com/fisher-lebo/ocd/modules/syndication"
)

type RDF struct {
	XMLName xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
	Channel Channel  `xml:"channel"`
	Items   []Item   `xml:"item"`
}

type Channel struct {
	About       string    `xml:"about,attr"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Image       Image     `xml:"image"`
	Items       []Li      `xml:"items>Seq>li"`
	TextInput   TextInput `xml:"textinput"`
	dublincore.DublinCore
	syndication.Syndication
}

type Item struct {
	About       string `xml:"about,attr"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	dublincore.DublinCore
}

type Image struct {
	About string `xml:"about,attr"`
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Link  string `xml:"link"`
	dublincore.DublinCore
}

type Li struct {
	Resource string `xml:"resource,attr"`
}

type TextInput struct {
	About       string `xml:"about,attr"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
	dublincore.DublinCore
}
