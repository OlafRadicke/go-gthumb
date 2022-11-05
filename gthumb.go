package gthumb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
)

type commentfile interface {
	NewCommentFile(path string)
	Load()
	Save()
	AddCategory(string)
	GetCaption()
	GetCategories()
	GetCommentVersion()
	GetNote()
	RemoveCategory(string)
	SetCaption()
	SetNote()
}

type CommentFile struct {
	FilePath string
	XML XMLComment
}

// NewCommentFile Inite an CommentFile and get it back.
// It's need an steing with the path to the xml comment file
// of gThumb.
func NewCommentFile(path string) (*CommentFile, error) {
	commentFile := CommentFile{}
	commentFile.FilePath = path
	err := commentFile.Load()
	if err != nil {
		return nil, fmt.Errorf("can't read comment file: %s", err)
	}
	return &commentFile, nil
}

// Load Loaded the conten of the comment file.
func (commentFile *CommentFile) Load() (error) {
	var err error
	xmlFile, err := os.Open(commentFile.FilePath)
	if err != nil {
		return fmt.Errorf("can't read comment file: %s", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	err = xml.Unmarshal(byteValue, &commentFile.XML)
	if err != nil {
		return fmt.Errorf("can't unmarshal comment file: %s", err)
	}
	return nil
}


// Save write this comment object back to the file in the .comment directory.
func (commentFile *CommentFile) Save() (error) {

	xmlString, err := xml.MarshalIndent(commentFile.XML, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal xml comment file: %s", err)
	}
	xmlFile, err := os.Create(commentFile.FilePath)
    if err != nil {
        return fmt.Errorf("can't create xml comment file: %s", err)
    }
	defer xmlFile.Close()
    _, err = xmlFile.WriteString(xml.Header + string(xmlString))
    if err != nil {
		return fmt.Errorf("can't write xml comment file: %s", err)
    }
	return nil
}


// GetCommentVersion Get back the version Categories xml file format.
func (commentFile *CommentFile) GetCommentVersion() (string) {
	return commentFile.XML.CommentVersion
}

// GetCaption Get back the value of Caption.
func (commentFile *CommentFile) GetCaption() (string) {
	return commentFile.XML.Caption
}

// SetCaption Set the value of Caption.
func (commentFile *CommentFile) SetCaption(newValue string) () {
	commentFile.XML.Caption = newValue
}

// GetNote Get back the value of Note.
func (commentFile *CommentFile) GetNote() (string) {
	return commentFile.XML.Note
}

// SetNote Set the value of Note.
func (commentFile *CommentFile) SetNote(newValue string) () {
	commentFile.XML.Note = newValue
}

// GetCategories Get back the list of the Categories
func (commentFile *CommentFile) GetCategories() ([]xmlCategory) {
	return commentFile.XML.Categories.CategoryList
}

// AddCategory Add a Category to the list of the Categories
func (commentFile *CommentFile) AddCategory(catValue string) () {
	newCat := xmlCategory{Value: catValue}
	commentFile.XML.Categories.CategoryList = append(commentFile.XML.Categories.CategoryList, newCat)
}

// RemoveCategory Remove a Category from the list of the Categories
func (commentFile *CommentFile) RemoveCategory(catValue string) () {
	var newList []xmlCategory

	for _, category := range commentFile.GetCategories() {
		if category.Value == catValue {
			continue;
		} else {
			newList = append(newList, category)
		}
	}
	commentFile.XML.Categories.CategoryList = newList
}