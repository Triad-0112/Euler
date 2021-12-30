package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Graduate struct {
	Success bool   `json:"success"`
	Result  Result `json:"result"`
}

type Result struct {
	Resource_id string    `json:"resource_id"`
	Fields      []Fields  `json:"fields"`
	Records     []Records `json:"records"`
}

type Fields struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type Records struct {
	Ide    int    `json:"_id"`
	Sex    string `json:"sex"`
	No     string `json:"no_of_graduates"`
	Course string `json:"type_of_course"`
	Year   string `json:"year"`
}

func main() {
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=660"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := Graduate{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	go func() {
		records := [][]string{}
		for y := 1993; y <= 2014; y++ {
			convert := strconv.Itoa(y)
			for i := 0; i < len(people1.Result.Records); i++ {
				if people1.Result.Records[i].Year == convert {
					temp := []string{
						strconv.Itoa(people1.Result.Records[i].Ide),
						people1.Result.Records[i].Sex,
						people1.Result.Records[i].No,
						people1.Result.Records[i].Course,
						people1.Result.Records[i].Year,
					}
					records = append(records, temp)
				}
			}
			filename := convert + ".csv"
			CreateFile(filename, records)
			records = records[:0]
		}
	}()
	fmt.Scanln()
}
func CreateFile(filename string, records [][]string) {
	f, err := os.Create(filename)

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}
