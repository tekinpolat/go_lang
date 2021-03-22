package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)


func main() {
	url := "https://kur.doviz.com/api/v5/converterItems"
	method := "GET"

	for i := 0; i < 10000; i++{
		currentTime := time.Now()
		client := &http.Client {}

		req, _ := http.NewRequest(method, url, nil)
	
		res, _ := client.Do(req)
	
		defer res.Body.Close()
	
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(string(body))
		fmt.Printf("%v %d. request\n", currentTime.Format("2006-01-02 13:04:05"),i)
	}

}