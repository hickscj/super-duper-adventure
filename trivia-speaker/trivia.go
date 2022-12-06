package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

// Medium article for reference
// https://medium.com/nerd-for-tech/your-first-golang-rest-api-client-287c8dc0961
// this can handle unknown json - []map[string]interface{}
type data struct {
	Results []Question `json:"results"`
}

func sayIt(thingToSay string) {
	cmd := exec.Command("say", thingToSay)

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func getQuestions() {
	data_obj := data{}
	url := "https://opentdb.com/api.php?amount=50&category=27&difficulty=easy&type=boolean"
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
	q1 := data_obj.Results[0]
	if q1.Type == "boolean" {
		sayIt("True or false: " + q1.Question)
	}
	time.Sleep(10 * time.Second)
	sayIt(q1.CorrectAnswer)
}

func main() {
	getQuestions()
}
