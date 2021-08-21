package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"github.com/atotto/clipboard"
)

// We are only interested in all types of images present in a directory
var allowedExtensions = []string{"jpg", "jpeg", "png"}

// Check whether a filename represents an image
func isAllowedFileType(fileName string) bool {
	valid := false
	for _, extension := range allowedExtensions {
		valid = valid || strings.HasSuffix(fileName, extension)
	}

	return valid
}

// Traverse the directory and extract the names of all file/folders
func readDirectoryContents(root string) ([]string, error) {
	var fileNames []string
	filesInfo, err := ioutil.ReadDir(root)

	if err != nil {
		return fileNames, err
	}

	for _, file := range filesInfo {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

// Filter all file names that represent images
func extractAllowedFileNames(fileNames []string) []string {
	var res []string

	for _, fileName := range fileNames {
		if isAllowedFileType(fileName) {
			res = append(res, fileName)
		}
	}

	return res
}

// Process the image list and customize the output
func extractCustomizedList(fileNames []string) string {
	var collatedMessage string
	collatedMessage += "\\begin{center}\n"
	for _, fileName := range fileNames {
		collatedMessage += "\\includegraphics{" + fileName + "}\n"
	}
	collatedMessage += "\\end{center}"

	return collatedMessage
}

func main() {
	// The first argument is the name of the program.
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory name")
		return
	}

	// Extract the working directory to enable relative paths
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	completeDirectoryName := currentWorkingDirectory + "/" + os.Args[1]

	// Extract the names of all files/folders in the desired directory
	fileNames, err := readDirectoryContents(completeDirectoryName)
	if err != nil {
		fmt.Println("There's no such directory: " + completeDirectoryName)
		return
	}

	// Filter the images
	allowedFileNames := extractAllowedFileNames(fileNames)

	if len(allowedFileNames) == 0 {
		fmt.Println("There are no images in the mentioned directory")
		return
	}

	// Collect the output in a string to copy to clipboard
	collatedMessage := extractCustomizedList(allowedFileNames)
	fmt.Println(collatedMessage)

	// Copy to clipboard
	clipboard.WriteAll(collatedMessage)

	// Success
	fmt.Println("\n**************************************************\n")
	fmt.Println("Text has been successfully copied to the clipboard")
	fmt.Println("\n**************************************************\n")
}
