package atom

import (
	"encoding/xml"
	"time"
)

type Feed struct {
	Metadata
	XMLName xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Source  Metadata `xml:"source"`
	Entries []Entry  `xml:"entry"`
}

type Metadata struct {
	Id           string     `xml:"id"`
	Title        string     `xml:"title"`
	Updated      time.Time  `xml:"updated"`
	Authors      []Person   `xml:"author"`
	Links        []Link     `xml:"link"`
	Categories   []Category `xml:"category"`
	Contributors []Person   `xml:"contributor"`
	Generator    Generator  `xml:"generator"`
	Icon         string     `xml:"icon"`
	Logo         string     `xml:"logo"`
	Rights       Text       `xml:"rights"`
	Subtitle     Text       `xml:"subtitle"`
}

type Entry struct {
	Id           string     `xml:"id"`
	Title        string     `xml:"title"`
	Updated      time.Time  `xml:"updated"`
	Authors      []Person   `xml:"author"`
	Content      Content    `xml:"content"`
	Links        []Link     `xml:"link"`
	Summary      Text       `xml:"summary"`
	Categories   []Category `xml:"category"`
	Contributors []Person   `xml:"contributor"`
	Published    time.Time  `xml:"published"`
	Source       Metadata   `xml:"source"`
	Rights       Text       `xml:"rights"`
}

type Person struct {
	Name  string `xml:"name"`
	Uri   string `xml:"uri"`
	Email string `xml:"email"`
}

type Link struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}

type Category struct {
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr"`
	Label  string `xml:"label,attr"`
}

type Generator struct {
	Uri     string `xml:"uri,attr"`
	Version string `xml:"version,attr"`
	Text    string `xml:",chardata"`
}

type Text struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

type Content struct {
	Type string `xml:"type,attr"`
	Src  string `xml:"src,attr"`
	Text string `xml:",chardata"`
}
