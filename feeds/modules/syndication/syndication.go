package syndication

import (
	"time"
)

type Syndication struct {
	UpdatePeriod    string    `xml:"http://purl.org/rss/1.0/modules/syndication updatePeriod"`
	UpdateFrequency int       `xml:"http://purl.org/rss/1.0/modules/syndication updateFrequency"`
	UpdateBase      time.Time `xml:"http://purl.org/rss/1.0/modules/syndication updateBase"`
}
