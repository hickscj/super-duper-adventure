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

type Questions struct {
	Results []Question `json:"results"`
}

func getQuestions() {
	var data []Questions
	url := "https://opentdb.com/api.php?amount=10&category=27&difficulty=easy&type=boolean"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
	jsonErr := json.Unmarshal([]byte(string(body)), &data)
	if err != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println("hey")
	for _, res := range data {
		fmt.Println(res.Results[0])
		// fmt.Printf("Question: %s", res)
	}
}

func main() {
	getQuestions()
}
