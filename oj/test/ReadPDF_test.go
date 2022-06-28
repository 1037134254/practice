package test

import (
	"fmt"
	"github.com/google/go-tika/tika"
	"golang.org/x/net/context"
	"log"
	"os"
	"testing"
)

func TestPdf(t *testing.T) {
	// Optionally pass a port as the second argument.
	f, err := os.Open("C:\\oj\\test\\1.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(f.Name())
	client := tika.NewClient(nil, "http://localhost:9998")
	body, err := client.Parse(context.Background(), f)
	fmt.Println(err)
	fmt.Println(body)
}
