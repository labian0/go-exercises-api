package main

import (
	"fmt"
	"os"
	"ra/ctf-api/internal/api"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading dotenv: " + err.Error())
		return
	}
	var apiURL string = os.Getenv("API_URL")
	if apiURL == "" {
		fmt.Println("error loading dotenv: API_URL is empty")
		return
	}
	var ex1AnswerRaw string = os.Getenv("EX1_ANSWER")
	if ex1AnswerRaw == "" {
		fmt.Println("error loading dotenv: EX1_ANSWER is empty")
		return
	}
	ex1Answer, err := strconv.Atoi(ex1AnswerRaw)
	if err != nil {
		fmt.Printf("error converting ex1AnswerRaw (%s) to int: %s", ex1AnswerRaw, err.Error())
		return
	}
	api.Init(apiURL, ex1Answer)
	err = api.Run()
	fmt.Println(err.Error())
}
