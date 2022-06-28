package test

import (
	"fmt"
	"os"
	"testing"
)

func TestOS(t *testing.T) {
	open, _ := os.Open("C:\\oj\\test\\copy.docx")
	bytes := make([]byte, 10)
	at, _ := open.ReadAt(bytes, 1024)
	fmt.Println(at)
	fmt.Println(string(bytes[0:at]))
	at1, _ := open.ReadAt(bytes, 1028)
	fmt.Println(at1)
	fmt.Println(string(bytes[:10]))
	at2, _ := open.ReadAt(bytes, 20)
	fmt.Println(at2)
	fmt.Println(string(bytes[:10]))
}
