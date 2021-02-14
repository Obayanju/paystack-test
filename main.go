package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const BASE_URL = "https://api.paystack.co"

var SECRET_KEY = os.Getenv("PAYSTACK_SECRET_KEY")

type CustomerInfo struct {
	Email  string `json:"email"`
	Amount string `json:"amount"`
}

func main() {
	client := &http.Client{}

	custInfo := CustomerInfo{
		Email:  "jerry@gmail.com",
		Amount: "1000",
	}

	b, err := json.Marshal(custInfo)
	if err != nil {
		fmt.Println("error: ", err)
	}

	endpoint := fmt.Sprintf("%s/transaction/initialize", BASE_URL)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(b))
	authorizeVal := fmt.Sprintf("Bearer %s", SECRET_KEY)
	req.Header.Add("Authorization", authorizeVal)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
