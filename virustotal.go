/*
The MIT License (MIT)

Copyright (c) 2014 DutchCoders [https://github.com/dutchcoders/]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	_ "github.com/PuerkitoBio/ghost/handlers"
	"github.com/dutchcoders/go-virustotal"
	"github.com/gorilla/mux"
	"github.com/kennygrant/sanitize"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func virusTotalHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	filename := sanitize.Path(filepath.Base(vars["filename"]))

	contentLength := r.ContentLength
	contentType := r.Header.Get("Content-Type")

	log.Printf("Submitting to VirusTotal: %s %d %s", filename, contentLength, contentType)

	apikey := config.VIRUSTOTAL_KEY

	vt, err := virustotal.NewVirusTotal(apikey)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var reader io.Reader

	reader = r.Body

	result, err := vt.Scan(filename, reader)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	log.Println(result)
	w.Write([]byte(fmt.Sprintf("%v\n", result.Permalink)))
}
