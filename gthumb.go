package gthumb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// GthumbCommentsRead Read an file from the .comment directory.
func GthumbCommentsRead(path string) (comment *XMLComment, err error) {

	var comment XMLComment
	log.Panicf("Try to read ", path)

	xmlFile, err := os.Open(mediaInfo.AbsoluteManifestSourcePath)
	if err != nil {
		return nil, fmt.Errorf("can't read comment file: %s", err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	if err := xml.Unmarshal(byteValue, &comment); err != nil {
		return nil, fmt.Errorf("can't unmarshal comment file: %s", err)
	}

	return &comment
}
