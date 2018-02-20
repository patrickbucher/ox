package main

import (
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
		fmt.Fprintf(os.Stderr, "GET %s: %s\n", searchURL, resp.Status)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
