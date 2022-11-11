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
	Rating         XmlRating     `xml:"rating"`
	Time           XmlTime       `xml:"time"`
	Categories     XmlCategories `xml:"categories"`
}

type XmlRating struct {
	Value string `xml:"value,attr"`
}

type XmlTime struct {
	Value string `xml:"value,attr"`
}

type XmlCategories struct {
	CategoryList []XmlCategory `xml:"category"`
	Value        string        `xml:"value,attr"`
}

type XmlCategory struct {
	Value string `xml:"value,attr"`
}

func NewXMLComment() *XMLComment {
	xmlComment := XMLComment{}
	xmlComment.XMLName = xml.Name{}
	xmlComment.CommentVersion = ""
	xmlComment.Caption = ""
	xmlComment.Note = ""
	xmlComment.Place = ""
	xmlComment.Rating = XmlRating{}
	xmlComment.Rating.Value = ""
	xmlComment.Time = XmlTime{}
	xmlComment.Time.Value = ""
	xmlComment.Categories = XmlCategories{}
	xmlComment.Categories.CategoryList = []XmlCategory{}
	xmlComment.Categories.Value = ""

	return &xmlComment
}