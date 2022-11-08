package gthumb

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileTree struct {
	// The list of found comment files
	ListOfCommentFiles []string
	// The list of found media files
    ListOfMediaFiles []string
	// The start point in the file tree
	CollectionPath string
}

// NewProgArguments create new instance of FileTree and get it back.
func NewFileTree(path string) *FileTree {
	fileTree := FileTree{}
	fileTree.CollectionPath = path
	return &fileTree
}

// addCommentFilePath Store are comment file path in a list
func (fileTree *FileTree)addCommentFilePath(path string){
	fileTree.ListOfCommentFiles = append(fileTree.ListOfCommentFiles, path)
}

// addMediaFilePath Store a pathe of an Madia file in a list
func (fileTree *FileTree)addMediaFilePath(path string){
	fileTree.ListOfMediaFiles = append(fileTree.ListOfMediaFiles, path)
}

// GoThroughCollection This function gos through the directory of Media
// files recursively. All comment files that were found are listed in
// the property list ListOfCommentFiles. And all Media Files are listed
// in the property list ListOfMediaFiles.
func (fileTree *FileTree) GoThroughCollection() error{
	err := filepath.Walk(fileTree.CollectionPath, fileTree.fileHandler)
	if err != nil {
		return fmt.Errorf("can't walk to: %s", err)
	}
	return nil
}


// fileHandler This function handel and prosessing the file operations.
func (fileTree *FileTree) fileHandler(searchPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	} else {
		if filepath.Ext(info.Name()) == ".xml" {
			fileDir := filepath.Dir(searchPath)
			fileDirParts := strings.Split(fileDir, "/")
			parentDir := fileDirParts[len(fileDirParts) - 1]
			if parentDir == ".comments" {
				fileTree.addCommentFilePath(searchPath)
			} else {
				return nil
			}

		} else {
			fileTree.addMediaFilePath(searchPath)
		}
		return nil
	}
}