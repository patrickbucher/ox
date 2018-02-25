package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"ox"
)

func main() {
	lang := flag.String("lang", "en", "Lookup Language")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage: ox [word]\n")
		os.Exit(1)
	}
	wordId := url.QueryEscape(args[0])
	searchURL := fmt.Sprintf("%s/entries/%s/%s",
		ox.APIBaseURL, *lang, wordId)
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create request: %v\n", err)
		os.Exit(1)
	}
	req.Header.Add("app_id", ox.ApplicationId)
	req.Header.Add("app_key", ox.ApplicationKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "GET %s: %v\n", searchURL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			fmt.Fprintf(os.Stderr, "'%s' was not found\n", wordId)
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "GET %s: %s\n", searchURL, resp.Status)
			os.Exit(1)
		}
	}
	buf := bytes.NewBufferString("")
	io.Copy(buf, resp.Body)
	var result ox.EntryResponse
	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(result.String())
}
