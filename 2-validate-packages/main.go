package main
import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
        "math/rand"
)
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
func main() {
	ifile := os.Args[1]
	lines, err := readLines("in/"+ifile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	endpoints, err := readLines("endpoints.txt")
        if err != nil {
                log.Fatalf("readLines: %s", err)  
        }
	outfile, fileErr := os.Create("out/"+ifile)
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}
	defer outfile.Close()
	errfile, fileErr := os.Create("err/"+ifile)
        if fileErr != nil {
                fmt.Println(fileErr)
                return
        }
	defer errfile.Close()
	slice := lines
	sliceLength := len(slice)
	maxNbConcurrentGoroutines := flag.Int("maxNbConcurrentGoroutines", 100, "the number of goroutines that are allowed to run concurrently")
	nbJobs := flag.Int("nbJobs", sliceLength, "the number of jobs that we need to do")
	flag.Parse()
	concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	for i := 0; i < *maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}
	done := make(chan bool)
	waitForAllJobs := make(chan bool)
	go func() {
		for i := 0; i < *nbJobs; i++ {
			<-done
			concurrentGoroutines <- struct{}{}
		}
		waitForAllJobs <- true
	}()
	for i := 0; i < *nbJobs; i++ {
		<-concurrentGoroutines
		go func(id int) {
			val := slice[id]
			n := rand.Int() % len(endpoints)
			client := &http.Client{}

			req, err := http.NewRequest("GET", "https://" + endpoints[n] + "/fireprox?id=" + val, nil)
			if err != nil {
				log.Println(err)
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}

			if resp.StatusCode == 200 {
				fmt.Fprintf(outfile, "%v\n", val)
			} else {
				fmt.Fprintf(errfile, "%[1]v : %[2]v\n", val, resp.StatusCode)
			}
			fmt.Println("The status code we got is:", resp.StatusCode, val)
			done <- true
		}(i)
	}
	<-waitForAllJobs
}
