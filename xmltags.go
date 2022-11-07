package gthumb

type XMLTags struct {
	TagsVersion string      `xml:"version,attr"`
	TagList    	[]xmlTag 	`xml:"tag"`
}

type xmlTag struct {
	Value string `xml:"value,attr"`
}

