package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"sync"
	"time"
)

func newRequest() *http.Request {
	var buf bytes.Buffer
	var contentType string

	{
		w := multipart.NewWriter(&buf)
		file, err := os.Open("/html/base.html")
		if err != nil {
			log.Fatal(`os.Open("/html/base.html"):`, err)
		}
		formfile, err := w.CreateFormFile("html", "/html/base.html")
		if err != nil {
			log.Fatal(`w.CreateFormFile("html", "/html/base.html")`, err)
		}
		if _, err := io.Copy(formfile, file); err != nil {
			log.Fatal("io.Copy(fileField, file)", err)
		}
		file.Close()
		contentType = w.FormDataContentType()
		w.Close()
	}

	req, err := http.NewRequest(http.MethodPost, "http://convert:8080/html-to-pdf", &buf)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)

	return req
}

func run(i int, wg *sync.WaitGroup) {
	fmt.Println("run", i)
	defer wg.Done()
	req := newRequest()
	start := time.Now()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	f, err := os.Create(fmt.Sprintf("/pdf/%d.pdf", i))
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(f, res.Body); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("started: %s\tduration: %s\n", start.Format("02 Jan 06 15:04"), time.Since(start))
}

func main() {
	n := flag.Int("n", 1, "the number of requests")
	flag.Parse()

	wg := new(sync.WaitGroup)
	wg.Add(*n)

	fmt.Printf("will be run %d requests\n", *n)
	for i := 0; i < *n; i++ {
		go run(i, wg)
	}

	wg.Wait()
	fmt.Println("done!")
}
