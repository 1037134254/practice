package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	// 管道
	pipe, err := cmd.StdinPipe()
	if err != nil {
		log.Printf("err:" + err.Error())
		return
	}
	io.WriteString(pipe, "23 11\n")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	println("Err", stderr.Bytes())
	fmt.Println(out.Bytes())
	println(out.String() == "34\n")
}
