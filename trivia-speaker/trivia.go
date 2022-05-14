package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Question struct {
	Category         string `json:"category"`
	Type             string `json:"type"`
	Difficulty       string `json:"difficulty"`
	Question         string `json:"question"`
	CorrectAnswer    string
	IncorrectAnswers []string
}

// Medium article for reference
// https://medium.com/nerd-for-tech/your-first-golang-rest-api-client-287c8dc0961
// this can handle unknown json - []map[string]interface{}
type data struct {
	Results []Question `json:"results"`
}

func getQuestions() {
	data_obj := data{}
	url := "https://opentdb.com/api.php?amount=10&category=27&difficulty=easy&type=boolean"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// fmt.Println(string(body))
	jsonErr := json.Unmarshal(body, &data_obj)
	if err != nil {
		log.Fatal(jsonErr)
	}
	for _, res := range data_obj.Results {
		fmt.Printf("Question: %s", res.Question)
	}
}

func main() {
	getQuestions()
}
