package dublincore

import (
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
