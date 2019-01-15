package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/pathutil"

	"github.com/bitrise-io/go-utils/fileutil"
)

func main() {
	fileName := os.Getenv("file_name")
	absFilePath, err := pathutil.AbsPath(fileName)
	if err != nil {
		log.Fatalf("Failed to determine absolute path of file (%s): %+v", fileName, err)
	}

	if err := fileutil.WriteStringToFile(absFilePath, os.Getenv("file_content")); err != nil {
		log.Fatalf("Failed to write into file: %+v", err)
	}

	// Output
	cmdLog, err := exec.Command("bitrise", "envman", "add", "--key", "GENERATED_TEXT_FILE_PATH", "--value", absFilePath).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to expose output with envman, error: %#v | output: %s", err, cmdLog)
		os.Exit(1)
	}

	os.Exit(0)
}
