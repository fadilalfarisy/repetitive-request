package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", RandomPost)

	server := new(http.Server)
	server.Addr = ":3000"

	fmt.Println("server started at localhost:3000")
	server.ListenAndServe()
}

func RandomPost(water int, wind int) {
	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}
	requestJSON, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", requestJSON)
	req.Header.Set("Content Type", "applicatoin/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
