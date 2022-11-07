package gthumb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
)

type commentsfile interface {
	NewCommentsFile(path string)
	Load()
	Save()
	AddCategory(string)
	GetCaption()
	GetCategories()
	GetCommentVersion()
	GetNote()
	GetPlace()
	GetRating()
	RemoveCategory(string)
	SetCaption()
	SetNote()
	SetPlacee(string)
	SetRating(string)
}

type CommentsFile struct {
	FilePath string
	XML XMLComment
}

// NewCommentsFile Inite an CommentsFile and get it back.
// It's need an steing with the path to the xml comment file
// of gThumb.
func NewCommentsFile(path string) (*CommentsFile, error) {
	commentFile := CommentsFile{}
	commentFile.FilePath = path
	err := commentFile.Load()
	if err != nil {
		return nil, fmt.Errorf("can't read comment file: %s", err)
	}
	return &commentFile, nil
}

// Load Loaded the values of the comment file.
func (commentFile *CommentsFile) Load() (error) {
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
func (commentFile *CommentsFile) Save() (error) {

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
func (commentFile *CommentsFile) GetCommentVersion() (string) {
	return commentFile.XML.CommentVersion
}

// GetCaption Get back the value of Caption.
func (commentFile *CommentsFile) GetCaption() (string) {
	return commentFile.XML.Caption
}

// SetCaption Set the value of Caption.
func (commentFile *CommentsFile) SetCaption(newValue string) () {
	commentFile.XML.Caption = newValue
}

// GetNote Get back the value of Note.
func (commentFile *CommentsFile) GetNote() (string) {
	return commentFile.XML.Note
}

// SetNote Set the value of Note.
func (commentFile *CommentsFile) SetNote(newValue string) () {
	commentFile.XML.Note = newValue
}

// GetPlace Get back the value of Place.
func (commentFile *CommentsFile) GetPlace() (string) {
	return commentFile.XML.Place
}

// SetPlace Set the value of Placee.
func (commentFile *CommentsFile) SetPlace(newValue string) () {
	commentFile.XML.Place = newValue
}

// GetRating Get back the value of Rating.
func (commentFile *CommentsFile) GetRating() (string) {
	return commentFile.XML.Rating.Value
}

// SetRating Set the value of Ratinge.
func (commentFile *CommentsFile) SetRating(newValue string) () {
	commentFile.XML.Rating.Value = newValue
}

// GetCategories Get back the list of the Categories
func (commentFile *CommentsFile) GetCategories() ([]xmlCategory) {
	return commentFile.XML.Categories.CategoryList
}

// AddCategory Add a Category to the list of the Categories
func (commentFile *CommentsFile) AddCategory(catValue string) () {
	newCat := xmlCategory{Value: catValue}
	commentFile.XML.Categories.CategoryList = append(commentFile.XML.Categories.CategoryList, newCat)
}

// RemoveCategory Remove a Category from the list of the Categories
func (commentFile *CommentsFile) RemoveCategory(catValue string) () {
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