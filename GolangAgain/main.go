package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type Digimon struct {
	Name  string `json:"name"`
	Image string `json:"img"`
	Level string `json:"level"`
}

const (
	works = 10
	file  = "digimon.txt"
)

var endpoint = "https://digimon-api.vercel.app/api/digimon/name/%s"
var wg sync.WaitGroup

func main() {

	srcFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer srcFile.Close()

	chanLine := make(chan string)
	go readingFile(srcFile, chanLine)

	wg.Add(1)
	go load(chanLine)

	wg.Wait()
}

func load(digimons chan string) {
	defer wg.Done()
	for i := 0; i < works; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for digimon := range digimons {
                                url := fmt.Sprintf(endpoint, digimon)
				fmt.Println(getDigimon(url))
			}
		}()
	}
}

func getDigimon(url string) Digimon {
	var digimons []Digimon

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("accept", "appliction/json")
	request.Header.Add("content-type", "appliction/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &digimons); err != nil {
		log.Fatal(err)
	}

	return digimons[0]

}

func readingFile(file *os.File, ch chan string) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		ch <- fmt.Sprint(err)
	}
	close(ch)
}
