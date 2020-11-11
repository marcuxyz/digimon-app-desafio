package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

// Digimon is a struct
type Digimon struct {
	Name  string `json:"name"`
	Img   string `json:"img"`
	Level string `json:"level"`
}

func getNames() []string {
	content, err := ioutil.ReadFile("digimon.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func request(url string) {
	defer wg.Done()

	var digimons []Digimon

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&digimons); err != nil {
		if err != nil {
			panic(err)
		}
	}
}

// tenho 2 threads trabalhando
// 1 - rodando a função main
// 2 - rodando a função request
func main() {
	now := time.Now()

	for _, name := range getNames() {
		go request("https://digimon-api.vercel.app/api/digimon/name/" + name)
		wg.Add(1)
	}

	wg.Wait()

	// sem goroutines durou 51.702802825s
	fmt.Println("Duration: ", time.Since(now))
}
