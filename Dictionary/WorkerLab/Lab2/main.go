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

//Data TYPE DONT CHANGE
type Data struct {
	Entries []Entries `json:"entries"`
}

type Entries struct {
	API         string `json:"API"`
	Description string `json:"Description"`
	HTTPS       bool   `json:"HTTPS"`
	Cors        string `json:"Cors"`
	Link        string `json:"Link"`
	Category    string `json:"Category"`
}

func main() {
	totalworker := flag.Int("concurrent_limit", 2, "Input total worker")
	dir := flag.String("output", "D:/Default/", "Destination Output file")
	flag.Parse()
	url := "https://api.publicapis.org/entries"
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
	GData := Data{}
	jsonErr := json.Unmarshal(body, &GData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	m := make(map[string][][]string)
	for i := range GData.Entries {
		for j := range GData.Entries {
			if GData.Entries[j].Category == GData.Entries[i].Category {
				temp := []string{
					GData.Entries[j].API,
					GData.Entries[j].Description,
					strconv.FormatBool(GData.Entries[j].HTTPS),
					GData.Entries[j].Cors,
					GData.Entries[j].Link,
					GData.Entries[j].Category,
				}
				m[GData.Entries[j].Category] = append(m[GData.Entries[j].Category], temp)
			}
		}
	}
	var wg sync.WaitGroup
	numJobs := len(m)
	jobs := make(chan [][]string, numJobs)
	results := make(chan [][]string, numJobs)
	for w := 1; w <= *totalworker; w++ {
		wg.Add(1)
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
		filename := j[0][5] + ".csv"
		//results <- j //to show The result of jobs, unnecessary
		fmt.Printf("\n%s : %s %s %s\n", workercolor("[Worker %d]", id), textcolor("Creating"), filenamecolor("%s.csv", j[0][5]), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
		CreateFile(dir, &filename, j)
		fmt.Printf("\n%s : %s %s %s %s %s\n", workercolor("[Worker %d]", id), textcolor("Finished creating"), filenamecolor("%s.csv", j[0][5]), textcolor("At"), directorycolor("%s", *dir), timecolor("%s", now.Format("15:04:05.999999999Z07:00")))
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
