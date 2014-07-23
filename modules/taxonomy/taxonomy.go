package taxonomy

type Taxonomy struct {
	Topics Topics `xml:"http://purl.org/rss/1.0/modules/taxonomy/ topics"`
}

type Topics struct {
	Lis []Li `xml:"Bag>li"`
}

type Li struct {
	Resource string `xml:"resource,attr"`
}
