package main

import (
	"log"
	"os/exec"

	// github.com/OlafRadicke/go-gthumb
	gt "github.com/OlafRadicke/go-gthumb"
)

func main() {
	testDir := "/tmp/go-gthumb-test"
	path := testDir + "/.comments/Hans_Makart-Graefin_Palffy.jpeg.xml"
	path2 := testDir + "/.comments/Branicka_Katharina.jpg.xml"
	testPreparation(testDir)
	modifieTest(path)
	modifieTest(path2)
}

// Create a copy of the directory with the test files
func testPreparation(testDir string) {
	srcDir := "./assets"

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

	//  Get...
	log.Printf("comment has comment version: %s\n", comment.GetCommentVersion())
	log.Printf("comment has caption: %s\n", comment.GetCaption())
	log.Printf("comment has node: %s\n", comment.GetNote())
	log.Printf("comment has place: %s\n", comment.GetPlace())
	log.Printf("comment has rating: %s\n", comment.GetRating())

	for index, category := range comment.GetCategories() {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}

	// Set, add and remove ...
	comment.SetCaption("An new Caption")
	comment.SetNote("New Note")
	comment.SetPlace("Set new place")
	comment.SetRating("16")
	comment.AddCategory("Apfel")
	comment.AddCategory("Handtuch")
	comment.RemoveCategory("Handtuch")

	// Double check...

	log.Printf("Check caption: %s\n", comment.GetCaption())
	log.Printf("Check node: %s\n", comment.GetNote())
	log.Printf("Check place: %s\n", comment.GetPlace())
	log.Printf("Check rating: %s\n", comment.GetRating())
	for index, category := range comment.GetCategories() {
		log.Printf("comment has category (%d): %s\n", index, category.Value)
	}
	// write back...

	err = comment.Save()
	if err != nil {
		log.Panicf("error writting comment file: %w", err)
	}
}
