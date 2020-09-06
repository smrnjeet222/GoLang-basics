package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Req is Request ....
func Req() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s ", robots)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		panic(error.Error())
	}
}
