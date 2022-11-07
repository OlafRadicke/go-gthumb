package main

import (
	"log"
	"os/exec"

	gt "github.com/OlafRadicke/go-gthumb"
)

func main() {
	testDir := "/tmp/go-gthumb-test"
	path := testDir + "/gthumb/tags.xml"
	testPreparation(testDir)
	modifieTest(path)
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
	tags, err := gt.NewTagsFile(path)
	if err != nil {
		log.Panicf("error to init object: %w", err)
	}

	// Print tgas
	for index, tagItem := range tags.GetTags() {
		log.Printf("Tag file has tag (%d): %s\n", index, tagItem.Value)
	}

	// Add new tag...
	tags.AddTag("new_tag_01")
	tags.AddTag("new_tag_02")
	tags.RemoveTag("new_tag_01")

	// check tag again

	for index, tagItem := range tags.GetTags() {
		log.Printf("Check has tag (%d): %s\n", index, tagItem.Value)
	}

	// save...

	err = tags.Save()
	if err != nil {
		log.Panicf("error writting comment file: %w", err)
	}

}