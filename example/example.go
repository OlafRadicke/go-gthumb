package main

import (
	"log"
	"os/exec"

	// github.com/OlafRadicke/go-gthumb
	gt "github.com/OlafRadicke/go-gthumb"
)

func main() {
	path := "./tmp_test/.comments/Hans_Makart-Graefin_Palffy.jpeg.xml"
	path2 := "./tmp_test/.comments/Branicka_Katharina.jpg.xml"
	testPreparation()
	modifieTest(path)
	modifieTest(path2)
}

// Create a copy of the directory with the test files
func testPreparation() {
	srcDir := "./assets"
	testDir := "./tmp_test"

	cmd := exec.Command("cp", "--recursive", srcDir, testDir)
	cmd.Run()
}

func modifieTest(path string) {
	// Read xml...
	log.Printf("Try to read file 1 %s", path)
	comment, err := gt.NewCommentFile(path)
	if err != nil {
		log.Panicf("error to init object: %w", err)
	}

	log.Printf("comment has comment version: %s\n", comment.XML.CommentVersion)
	log.Printf("comment has caption: %s\n", comment.XML.Caption)
	log.Printf("comment has node: %s\n", comment.XML.Note)
	log.Printf("comment has place: %s\n", comment.XML.Place)
	log.Printf("comment has rating: %s\n", comment.XML.Rating.Value)

	for index, category := range comment.GetCategories() {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}

	comment.AddCategory("Handtuch")

	for index, category := range comment.GetCategories() {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}

	comment.RemoveCategory("Handtuch")

	for index, category := range comment.GetCategories() {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}

	log.Printf("Check node: %s\n", comment.XML.Note)
	log.Printf("Check rating: %s\n", comment.XML.Rating.Value)
	comment.XML.Note = "Boote in einer Halle"
	comment.XML.Rating.Value = "16"
	log.Printf("Check node change: %s\n", comment.XML.Note)
	log.Printf("Check rating change: %s\n", comment.XML.Rating.Value)
	// write back...

	err = comment.Save()
	if err != nil {
		log.Panicf("error writting comment file: %w", err)
	}
}
