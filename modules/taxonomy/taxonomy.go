package taxonomy

type Taxonomy struct {
	Topics Bag `xml:"http://purl.org/rss/1.0/modules/taxonomy/ topics"`
}

type Bag struct {
	Lis []Li `xml:"Bag>li"`
}

type Li struct {
	Resource string `xml:"resource,attr"`
}
