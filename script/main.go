package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func run(cmd *exec.Cmd) {
	start := time.Now()
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("started: %s\tduration: %s\n", start.Format("02 Jan 06 15:04"), time.Since(start))
}

func main() {
	cmd := exec.Command("node", "src/convert.js", "/html/base.html", "/pdf/base.pdf")
	cmd.Dir, cmd.Stdout, cmd.Stderr = "/converter", os.Stdout, os.Stderr
	run(cmd)
	log.Println("done!")
}
