package test

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	open, _ := os.Open("C:\\oj\\test\\copy.docx")
	name := open.Name()

	fileInfo, _ := open.Stat()
	fmt.Println(fileInfo.ModTime())
	fmt.Println(name)
}
