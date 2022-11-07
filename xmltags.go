package gthumb

import (
	"encoding/xml"
)

type XMLTags struct {
	TagsVersion string      `xml:"version,attr"`
	Tags     	[]xmlTag 	`xml:"tags"`
}


type xmlTag struct {
	Value string `xml:"value,attr"`
}

