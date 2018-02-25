package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"ox"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [word]")
		os.Exit(1)
	}
	wordId := url.QueryEscape(os.Args[1])
	searchURL := ox.APIBaseURL + "/entries/en/" + wordId
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
	fmt.Printf("%s\n", result.String())
}
