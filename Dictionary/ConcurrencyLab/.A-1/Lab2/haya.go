package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

//TESTING WORKERPOOL
type Jobs struct {
	Filename string
	Data     [][][]string
}
type Worker struct {
	ID   string
	Dir  string
	Jobs string
	wLog *log.Logger
}

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
	//check := flag.String("grab", "", "Checking If user know what he do")
	//dir := flag.String("output", "", "Destination Output file")
	//flag.Parse()
	var wg sync.WaitGroup
	url := "https://data.gov.sg/api/action/datastore_search?resource_id=eb8b932c-503c-41e7-b513-114cffbe2338&limit=660"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
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

	datap := Graduate{}
	jsonErr := json.Unmarshal(body, &datap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	ch := make(chan [][]string, 1000)
	m := make(map[int][][]string)
	//records := [][]string{}
	wg.Add(1)
	go func() {
		for y := 1993; y <= 2014; y++ {
			convert := strconv.Itoa(y)

			for i := range datap.Result.Records {
				//for i := 0; i < len(datap.Result.Records); i++ {
				if datap.Result.Records[i].Year == convert {
					temp := []string{
						strconv.Itoa(datap.Result.Records[i].Ide),
						datap.Result.Records[i].Sex,
						datap.Result.Records[i].No,
						datap.Result.Records[i].Course,
						datap.Result.Records[i].Year,
					}
					m[y] = append(m[y], temp)
				}
			}

			//filename := convert + ".csv"
			//wg.Add(1)
			//go CreateFile(dir, filename, m[y], &wg)
		}
		wg.Done()
	}()
	wg.Add(1)
	go Workergrab(m, ch, &wg)
	wg.Wait()
	//total_worker := 2
	/*
		for i := 1; i <= total_worker; i++ {
			wg.Add(1)
			myWorker := Worker{}
			myWorker.ID = "ID= " + strconv.Itoa(i) + "\n"
			myWorker.wLog = log.New(os.Stderr, myWorker.ID, 1)
			go func(w *Worker) {

			}(&myWorker)
			wg.Done()
		}
		wg.Wait()
	*/
}
func Workergrab(m map[int][][]string, ch chan [][]string, wg *sync.WaitGroup) {
	for y := 1993; y <= 2014; y++ {
		ch <- m[y]
	}
	wg.Done()
}
func CreateFile(dir *string, filename string, a [][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	filepath, err := filepath.Abs(*dir + filename)
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
