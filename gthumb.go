package gthumb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// GthumbCommentsRead Read an file from the .comment directory.
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
