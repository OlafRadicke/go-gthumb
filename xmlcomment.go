package gthumb

import (
	"encoding/xml"
)

type XMLComment struct {
	XMLName        xml.Name      `xml:"comment"`
	Caption        string        `xml:"caption"`
	Categories     xmlCategories `xml:"categories"`
	CommentVersion string        `xml:"version,attr"`
	Note           string        `xml:"note"`
	Place          string        `xml:"place"`
	Rating         xmlRating     `xml:"rating"`
}

type xmlCategories struct {
	CategoryList []xmlCategory `xml:"category"`
	Value        string        `xml:"value,attr"`
}

type xmlCategory struct {
	Value string `xml:"value,attr"`
}

type xmlRating struct {
	Value string `xml:"value,attr"`
}
