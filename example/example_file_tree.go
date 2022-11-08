package main

import (
	"log"
	"os/exec"

	gt "github.com/OlafRadicke/go-gthumb"
)

func main() {
	testDir := "/tmp/go-gthumb-test"
	testPreparation(testDir)
	walkInTree(testDir)
}

// Create a copy of the directory with the test files
func testPreparation(testDir string) {
	srcDir := "./assets"
	cmd := exec.Command("cp", "--recursive", srcDir, testDir)
	cmd.Run()
}

func walkInTree(path string) {
	// Read xml...
	log.Printf("Start walk in %s", path)

	fileTree := gt.NewFileTree(path)
	fileTree.GoThroughCollection()

	for index, Item := range fileTree.ListOfCommentFiles {
		log.Printf("Comment files (%d): %s\n", index, Item)
	}

	for index, Item := range fileTree.ListOfMediaFiles {
		log.Printf("Media files (%d): %s\n", index, Item)
	}

}