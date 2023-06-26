package main
import (
        "bufio"
        "flag"
        "fmt"
        "log"
        "os"
	"encoding/json"
	"github.com/payatu/google-play-scraper/pkg/app"
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

			x := 0
			y := 0
			a := app.New(val, app.Options{
				Country:  "us",
				Language: "us",
			})
			err := a.LoadDetails()
			if err != nil {
				x = 1
				fmt.Fprintf(errfile, "LoadDetails : %v\n", a.ID)
			}
			err = a.LoadPermissions()
			if err != nil {
				y=1
				fmt.Fprintf(errfile, "LoadPermissions : %v\n", a.ID)
			}

			if x == 0 && y == 0 {
				b, err := json.Marshal(a)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Fprintf(outfile, "%v\n", string(b))
			}
                        done <- true
                }(i)
        }
        <-waitForAllJobs
}
