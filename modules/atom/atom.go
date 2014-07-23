package atom

type Atom struct {
	Link Link `xml:"http://www.w3.org/2005/Atom link"`
}

type Link struct {
	Href     string `xml:"href,attr"`
	Rel      string `xml:"rel,attr"`
	Type     string `xml:"type,attr"`
	Hreflang string `xml:"hreflang,attr"`
	Title    string `xml:"title,attr"`
	Length   string `xml:"length,attr"`
}
