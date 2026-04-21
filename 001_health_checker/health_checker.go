package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)


var httpClient = &http.Client{
		Timeout: 5 * time.Second,
}



func checkURL(url string) (int, error){
	if!strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://"){
		url = "https://" + url	
	}


	response, err := httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	return response.StatusCode, nil
}

func main(){
	fmt.Println("Starting programm...")
	fmt.Println("Reading URLs...")
	
	data, err := os.Open("urls.txt")	

	if(err != nil){
		fmt.Printf("Error: file not found %v\n", err)
		return 
	}

	defer data.Close()

	var urlsList []string
	var linesCounter int = 0

	scanner := bufio.NewScanner(data)

	for scanner.Scan(){
		linesCounter += 1

		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#"){
			continue
		}

		urlsList = append(urlsList, line)
	}

	for _, url := range urlsList {
		responseCode, err := checkURL((url))
		if(err != nil){
			fmt.Printf("Failed to check %v\n", err)
		} else {
			fmt.Printf("Response code for %s is %d\n", url, responseCode)
		}
		
		
	}



	fmt.Printf("Processed %d lines\n", linesCounter)



	

}	






