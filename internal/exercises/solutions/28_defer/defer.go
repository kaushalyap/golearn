package deferr

import (
	"fmt"
	"os"
	"sync"
)

var (
	fileInstance *os.File
	once         sync.Once
)

func WriteToFile(*os.File) {
	f := CreateFile()
	defer f.Close()
	fmt.Fprintln(f, "data")
}

// Create a singleton of a temporary file
func CreateFile() *os.File {
	once.Do(func() {
		f, err := os.CreateTemp("", "tmpfile")
		if err != nil {
			panic(err)
		}
		fileInstance = f
	})
	return fileInstance
}
