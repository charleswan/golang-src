package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	arr := [...]string{
		"https://tnx.nl/ip",
		"https://ident.me/",
		"https://icanhazip.com/",

		"http://whatismyip.akamai.com/",
		"http://myip.dnsomatic.com/",
		"http://tnx.nl/ip",
		"http://ident.me/",
		"http://icanhazip.com/"}

	for _, v := range arr {
		if err := httpGet(v); err == nil {
			break
		}
	}
}

func httpGet(theurl string) error {
	resp, err := http.Get(theurl)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(body))
	return nil
}
