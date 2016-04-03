package stats

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func fileCount() int {
	count, err := exec.Command("/bin/sh", "-c", getCommand()).Output()
	if err != nil {
		log.Printf("Error while retrieving file count %s\n", err.Error())
		return 0
	}
	return bytes.Count(count, []byte(`\n`))
}

func getCommand() string {
	return fmt.Sprintf("lsof -p %v", os.Getpid())
}
