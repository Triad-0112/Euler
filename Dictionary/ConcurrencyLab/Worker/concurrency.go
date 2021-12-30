package main

import (
	"/workerpool"
	"encoding/csv"
	"encoding/json"
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

	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	// Start Worker Pool.
	totalWorker := 5
	wp := workerpool.NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 2
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(1 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d", res.id, res.value)
	}

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

	records := [][]string{}

	for i := 0; i < len(people1.Result.Records); i++ {
		//fmt.Println("User Name: " + people1.Result.Records[i].Year)
		// records[i][0] = strconv.Itoa(people1.Result.Records[i].Ide)
		// records[i][1] = people1.Result.Records[i].Sex
		// records[i][2] = people1.Result.Records[i].No
		// records[i][3] = people1.Result.Records[i].Course
		// records[i][4] = people1.Result.Records[i].Year
		temp := []string{
			strconv.Itoa(people1.Result.Records[i].Ide),
			people1.Result.Records[i].Sex,
			people1.Result.Records[i].No,
			people1.Result.Records[i].Course,
			people1.Result.Records[i].Year,
		}
		records = append(records, temp)
	}

	f, err := os.Create("users.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}

}
