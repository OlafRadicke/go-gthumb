package gthumb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type TagsFile struct {
	FilePath string
	XML XMLTags
}

// NewTagsFile Inite an TagsFile and get it back.
// It's need an steing with the path to the xml tags file
// of gThumb.
func NewTagsFile(path string) (*TagsFile, error) {
	tagsFile := TagsFile{}
	tagsFile.FilePath = path
	err := tagsFile.Load()
	if err != nil {
		return nil, fmt.Errorf("can't read tags file: %s", err)
	}
	return &tagsFile, nil
}

// Load Loaded the valus of the tags file.
func (tagsFile *TagsFile) Load() (error) {
	var err error
	xmlFile, err := os.Open(tagsFile.FilePath)
	if err != nil {
		return fmt.Errorf("can't read tags file: %s", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &tagsFile.XML)
	if err != nil {
		return fmt.Errorf("can't unmarshal tags file: %s", err)
	}
	return nil
}


// Save write this tags object back to the file.
func (tagsFile *TagsFile) Save() (error) {

	xmlString, err := xml.MarshalIndent(tagsFile.XML, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal xml tags file: %s", err)
	}
	xmlFile, err := os.Create(tagsFile.FilePath)
    if err != nil {
        return fmt.Errorf("can't create xml tags file: %s", err)
    }
	defer xmlFile.Close()
    _, err = xmlFile.WriteString(xml.Header + string(xmlString))
    if err != nil {
		return fmt.Errorf("can't write xml tags file: %s", err)
    }
	return nil
}

// GetTags Get back the list of the tags
func (tagsFile *TagsFile) GetTags() ([]xmlTag) {
	return tagsFile.XML.TagList
}

// AddTag Add a tag to the list of the tags
func (tagsFile *TagsFile) AddTag(tagValue string) () {
	newTag := xmlTag{Value: tagValue}
	tagsFile.XML.TagList = append(tagsFile.XML.TagList, newTag)
}

// RemoveTag Remove a tag from the list of the tags
func (tagsFile *TagsFile) RemoveTag(catValue string) () {
	var newList []xmlTag

	for _, tagItem := range tagsFile.GetTags() {
		if tagItem.Value == catValue {
			continue;
		} else {
			newList = append(newList, tagItem)
		}
	}
	tagsFile.XML.TagList = newList
}