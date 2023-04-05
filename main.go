package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	doEvery(15000*time.Millisecond, helloworld)
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld(t time.Time) {
	rand.Seed(time.Now().UnixNano())

	randomWater := rand.Intn(100)
	randomWind := rand.Intn(100)

	data := map[string]interface{}{
		"water": randomWater,
		"wind":  randomWind,
	}

	var statusWater string
	var statusWind string

	if 5 >= randomWater {
		statusWater = "aman"
	} else if randomWater >= 6 && randomWater <= 8 {
		statusWater = "siaga"
	} else {
		statusWater = "bahaya"
	}

	if 6 >= randomWind {
		statusWind = "aman"
	} else if randomWind >= 7 && randomWind <= 15 {
		statusWind = "siaga"
	} else {
		statusWind = "bahaya"
	}

	requestJSON, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-type", "application/json")
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

	fmt.Println(string(body))
	fmt.Printf("status water : %s \n", statusWater)
	fmt.Printf("status wind : %s \n", statusWind)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", ActionRandom)

// 	server := new(http.Server)
// 	server.Addr = ":3000"

// 	fmt.Println("server started at localhost:3000")
// 	server.ListenAndServe()
// }

// func ActionRandom(w http.ResponseWriter, r *http.Request) {
// 	if !AllowOnlyGET(w, r) {
// 		return
// 	}

// 	if id := r.URL.Query().Get("userID"); id != "" {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(res)
// 		return
// 	}

// 	OutputJSON(w, GetStudents())
// }

// func OutputJSON(w http.ResponseWriter, o interface{}) {
// 	res, err := json.Marshal(o)
// 	if err != nil {
// 		w.Write(([]byte(err.Error())))
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }
