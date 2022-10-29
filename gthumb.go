package gthumb

import (
	"encoding/xml"
	"fmt"
	"log"
	"io/ioutil"
	"os"
)

// GthumbCommentsRead read an file from the .comment directory.
// Returned an Object with the information of the readed file.
func GthumbCommentsRead(path string) (*XMLComment, error) {
	var comment *XMLComment

	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't read comment file: %s", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	if err := xml.Unmarshal(byteValue, &comment); err != nil {
		return nil, fmt.Errorf("can't unmarshal comment file: %s", err)
	}
	return comment, nil
}


// GthumbCommentsRead write an comment file to the .comment directory.
func GthumbCommentsWrite(path string, comment *XMLComment) (error) {

	xmlString, err := xml.MarshalIndent(comment, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal xml comment file: %s", err)
	}
	xmlFile, err := os.Create(path)
    if err != nil {
        return fmt.Errorf("can't create xml comment file: %s", err)
    }
	defer xmlFile.Close()
	log.Printf(string(xmlString))
    _, err = xmlFile.WriteString(xml.Header + string(xmlString))
    if err != nil {
		return fmt.Errorf("can't write xml comment file: %s", err)
    }
	return nil
}