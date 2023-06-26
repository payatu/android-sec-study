package main

  import (
  	"fmt"
  	"strings"
	"bufio"
	"os"
	"io/ioutil"
  	xj "github.com/basgys/goxml2json"
  )

  func main() {

        readFile, err := os.Open("manifest.txt")
	if err != nil {
           fmt.Println(err)
	}
	defer readFile.Close()

	f, err := os.Create("manifest.json")
        if err != nil {
           fmt.Println(err)
        }
        defer f.Close()

	fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)

        for fileScanner.Scan() {

		fileContent, err := ioutil.ReadFile(fileScanner.Text())
	   	if err != nil {
	       	   fmt.Println(err)
	        }
		text := string(fileContent)
	  	xml := strings.NewReader(text)
	  	json, err := xj.Convert(xml)
	  	if err != nil {
	  		panic("That's embarrassing...")
		}
		fmt.Fprintf(f,"%v",json.String())
	}
  }
