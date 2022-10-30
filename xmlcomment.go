package gthumb

import (
	"encoding/xml"
)

type XMLComment struct {
	XMLName        xml.Name      `xml:"comment"`
	CommentVersion string        `xml:"version,attr"`
	Caption        string        `xml:"caption"`
	Note           string        `xml:"note"`
	Place          string        `xml:"place"`
	Rating         xmlRating     `xml:"rating"`
	Time           xmlTime       `xml:"time"`
	Categories     xmlCategories `xml:"categories"`
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

type xmlTime struct {
	Value string `xml:"value,attr"`
}