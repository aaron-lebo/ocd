package media

import (
	"time"
)

type Media struct {
	Title      Title       `xml:"http://search.yahoo.com/mrss/ title"`
	Thumbnails []Thumbnail `xml:"http://search.yahoo.com/mrss/ thumbnail"`
}

type Title struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

type Thumbnail struct {
	Url    string    `xml:"url,attr"`
	Height int       `xml:"height,attr"`
	Width  int       `xml:"width,attr"`
	Time   time.Time `xml:"time,attr"`
}
