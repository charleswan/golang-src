package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {

	resp, err := http.Get("https://golang.org")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// print out
	fmt.Println(os.Stdout, string(htmlData)) //<-- here !

	// use Regular Expression to search for keyword
	// for example
	verified, err := regexp.MatchString("VERIFIED", string(htmlData))

	//if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
}
