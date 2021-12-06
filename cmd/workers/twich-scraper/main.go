package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//url := "https://twitter.com/eloopcarsharing"
	//url := "https://www.linkedin.com/in/frederic-nachbauer"
	url := "https://instagram.com/utopiantravel"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
