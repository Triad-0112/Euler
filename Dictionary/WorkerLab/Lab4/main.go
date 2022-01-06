package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

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

func main() {
	start := time.Now()
	totalworker := flag.Int("concurrent_limit", 2, "Input total worker")
	dir := flag.String("output", "D:/Default/", "Destination Output file")
	flag.Parse()
	var wg sync.WaitGroup
	m := make(map[string][][]string)
	wg.Add(*totalworker)
	for w := 1; w <= *totalworker; w++ {
		go workerdata(w, m, &wg)
	}
	wg.Wait()
	//filename := convert + ".csv"
	//wg.Add(1)
	//go CreateFile(dir, filename, m[convert], &wg)

	numJobs := len(m)
	jobs := make(chan [][]string, numJobs)
	results := make(chan [][]string, numJobs)
	wg.Add(*totalworker)
	for w := 1; w <= *totalworker; w++ {
		go worker(w, jobs, results, dir, &wg)
	}
	for _, job := range m {
		jobs <- job
	}
	close(jobs)
	// Show each of jobs result
	/*for a := 1; a <= numJobs; a++ {
	//fmt.Println(<-results)
	}
	*/
	wg.Wait()
	fmt.Println(time.Since(start))
}
func workerdata(id int, m map[string][][]string, wg *sync.WaitGroup) map[string][][]string {
	textcolor := color.New(color.FgHiWhite, color.Bold).SprintfFunc()
	workercolor := color.New(color.FgHiCyan, color.Bold).SprintfFunc()
	filenamecolor := color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
	timecolor := color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
	now := time.Now()
	fmt.Printf("\n%s : %s %s\n", workercolor("[Worker %d]", id), textcolor("Starting to work at"), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=660"
	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
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
	defer wg.Done()
	people1 := Graduate{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	for y := 1993; y <= 2014; y++ {
		convert := strconv.Itoa(y)
		if m[convert] == nil {
			fmt.Printf("\n%s : %s%s %s\n", workercolor("[Worker %d]", id), textcolor("Processing Data-"), filenamecolor("%s", convert), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
			for i := range people1.Result.Records {
				//for i := 0; i < len(people1.Result.Records); i++ {
				if people1.Result.Records[i].Year == convert {
					temp := []string{
						strconv.Itoa(people1.Result.Records[i].Ide),
						people1.Result.Records[i].Sex,
						people1.Result.Records[i].No,
						people1.Result.Records[i].Course,
						people1.Result.Records[i].Year,
					}
					m[convert] = append(m[convert], temp)
				}
			}
			fmt.Printf("\n%s : %s %s %s\n", workercolor("[Worker %d]", id), textcolor("Done Processing-"), filenamecolor("%s", convert), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
		}
	}
	fmt.Printf("\n%s : %s %s\n", workercolor("[Worker %d]", id), textcolor("Finished work at"), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
	return m
}

func worker(id int, jobs <-chan [][]string, results chan<- [][]string, dir *string, wg *sync.WaitGroup) {
	defer wg.Done()
	textcolor := color.New(color.FgHiWhite, color.Bold).SprintfFunc()
	workercolor := color.New(color.FgHiCyan, color.Bold).SprintfFunc()
	directorycolor := color.New(color.FgHiYellow, color.Bold, color.Italic).SprintfFunc()
	filenamecolor := color.New(color.FgHiGreen, color.BlinkRapid, color.Bold).SprintfFunc()
	timecolor := color.New(color.FgHiMagenta, color.Bold).SprintfFunc()
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

/*
func worker(id int, j chan Jobs) {
	for jobs := range
}
*/
