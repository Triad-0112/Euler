package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	//"flag"

	"net/http"

	"github.com/fatih/color"
)

//TESTING WORKERPOOL
/*
type Jobs struct {
	Filename string
	Data     [][][]string
}

type Worker struct {
	ID   string
	Dir  string
	Jobs string
}
*/

//Data TYPE DONT CHANGE
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

const baseurl = "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&q="

var textcolor = color.New(color.FgHiWhite, color.Bold).SprintfFunc()
var mainworkercolor = color.New(color.FgHiRed, color.Bold).SprintFunc()
var workercolor = color.New(color.FgHiCyan, color.Bold).SprintfFunc()
var directorycolor = color.New(color.FgHiYellow, color.Bold, color.Italic).SprintfFunc()
var filenamecolor = color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
var timecolor = color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
var now = time.Now()

func main() {
	totalworker := flag.Int("concurrent_limit", 2, "Input total worker")
	dir := flag.String("output", "D:/Default/", "Destination Output file")
	flag.Parse()
	//datamap := make(map[string][][]string)
	var wg sync.WaitGroup
	ch := make(chan [][]string, 22)
	m := make(map[string][][]string)
	wg.Add(*totalworker)
	go workerfetch(m, ch, &wg, 1, 1993, 2014) // fetch Function Concurreny
	wg.Wait()
	for w := 2; w <= *totalworker; w++ {
		go worker(w, ch, dir, &wg)
	}
	close(ch)
}

func workerfetch(m map[string][][]string, ch chan [][]string, wg *sync.WaitGroup, id, yearinit, yearend int) {
	defer wg.Done()
	for i := yearinit; i <= yearend; i++ {
		url := baseurl + strconv.Itoa(i)
		spaceClient := http.Client{
			Timeout: time.Second * 10,
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
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
		record := Graduate{}
		jsonErr := json.Unmarshal(body, &record)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		convert := strconv.Itoa(i)
		for i := range record.Result.Records {
			temp := []string{
				strconv.Itoa(record.Result.Records[i].Ide),
				record.Result.Records[i].Sex,
				record.Result.Records[i].Course,
				record.Result.Records[i].Year,
			}
			m[convert] = append(m[convert], temp)
		}
		ch <- m[convert]
	}
}

func worker(id int, jobs <-chan [][]string, dir *string, wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()
	fmt.Printf("\n%s : %s %s\n", workercolor("[Worker %d]", id), textcolor("Starting to work at"), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
	for j := range jobs {
		filename := j[0][4] + ".csv"
		//results <- j //to show The result of jobs, unnecessary
		fmt.Printf("\n%s : %s %s %s\n", workercolor("[Worker %d]", id), textcolor("Creating"), filenamecolor("%s.csv", j[0][4]), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
		CreateFile(dir, &filename, j)
		fmt.Printf("\n%s : %s %s %s %s %s\n", workercolor("[Worker %d]", id), textcolor("Finished creating"), filenamecolor("%s.csv", j[0][4]), textcolor("At"), directorycolor("%s", *dir), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
	}
	fmt.Printf("\n%s : %s %s\n", workercolor("[Worker %d]", id), textcolor("Finished work at"), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
}
func CreateFile(dir *string, filename *string, a [][]string) {
	filepath, err := filepath.Abs(*dir + *filename)
	if err != nil {
		log.Fatalln("Invalid path")
	}
	f, err := os.Create(filepath)
	if err != nil {

		log.Fatalln("failed to open file", err)
	}
	//value := <-records
	w := csv.NewWriter(f)
	err = w.WriteAll(a) // calls Flush internally
	if err != nil {
		log.Fatal(err)
	}
}
