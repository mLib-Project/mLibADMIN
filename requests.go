package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var pageURL string = "https://mlib.mslotwinski.eu/api"

// var pageURL string = "http://localhost/api"

func get() {
	resp, err := http.Get(pageURL)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	fmt.Println(sb)
}

func post() {
	for i := 0; i <= len(books)-1; i++ {
		pushBook(books[i])
	}
}

func pushBook(book book) {
	postBody, _ := json.Marshal(map[string]string{
		"category":    categories[book.category].name,
		"family":      strconv.Itoa(categories[book.category].family),
		"subcategory": book.subcategory,
		"name":        book.name,
		"author":      book.author,
		"source":      book.source,
		"ID":          makeID(10),
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(pageURL, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

}

func delete() {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", pageURL+"/"+passwordCode, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}
