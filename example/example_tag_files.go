package main

import (
	"log"
	"os/exec"

	// github.com/OlafRadicke/go-gthumb
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

	for index, tagItem := range tags.GetTages() {
		log.Printf("comment has category (%d): %s\n", index, tagItem.Value)
	}

}