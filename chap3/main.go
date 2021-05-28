package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	log.Println(resp.Header.Get("Content-Length"))
	log.Println(resp.StatusCode)
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//		panic(err)
	//	}

	///##log.Println(string(body))
}
