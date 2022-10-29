package main

import (
	"log"
	"os"

	// github.com/OlafRadicke/go-gthumb
	gt "github.com/OlafRadicke/go-gthumb"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Please provide the path to the file to analyze")
		os.Exit(1)
	}
	path := os.Args[1]

	// Reat xml...
	log.Printf("Try to read %s", path)
	comment, err := gt.CommentsRead(path)
	if err != nil {
		log.Panicf("error parsing comment file: %w", err)
	}
	log.Printf("comment has comment version: %s\n", comment.CommentVersion)
	log.Printf("comment has caption: %s\n", comment.Caption)
	log.Printf("comment has node: %s\n", comment.Note)
	log.Printf("comment has place: %s\n", comment.Place)
	log.Printf("comment has rating: %s\n", comment.Rating.Value)

	for index, category := range comment.Categories.CategoryList {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}

	// write back...

	err = gt.CommentsWrite("./test-output.xml", comment)
	if err != nil {
		log.Panicf("error writting comment file: %w", err)
	}
}
