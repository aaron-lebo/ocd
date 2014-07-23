package taxonomy

type Taxonomy struct {
	Topics []Li `xml:"topics>Bag>li"`
}

type Li struct {
	Resource string `xml:"resource,attr"`
}
