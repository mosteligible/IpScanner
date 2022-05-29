package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"time"
)

type ipAddr struct {
	Ip string `json:"ip"`
}

type slackPayload struct {
	Message string `json:"text"`
}

func getIp() ipAddr {
	url := "https://api.ipify.org?format=json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request!")
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	ipAddress := ipAddr{}
	jsonErr := json.Unmarshal(body, &ipAddress)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return ipAddress
}

func sendSlackMsg(newIpAddr ipAddr) {
	webhookURL := os.Getenv("SLACK_WEBHOOK")
	payload := slackPayload {
		Message: newIpAddr.Ip,
	}
	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData) )

	// An error is returned if something goes wrong
	if err != nil {
		panic(err)
	}
	//Need to close the response stream, once response is read.
	//Hence defer close. It will automatically take care of it.
	defer resp.Body.Close()

	//Check response code, if New user is created then read response.
	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			panic(err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

	} else {
		//The status is not Created. print the error.
		fmt.Println("Get failed with error: ", resp.Status)
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Environment Variables NOT SET")
	}
	currentIp := ""
	for {
		mostRecentIp := getIp()
		fmt.Println(time.Now(), ": Ip Detection Request Sent. Most recent Ip:", mostRecentIp.Ip)
		if mostRecentIp.Ip != currentIp {
			fmt.Printf("Ip Address change detected, changed from < %s > to < %s >\n", currentIp, mostRecentIp.Ip)
			currentIp = mostRecentIp.Ip
			sendSlackMsg(mostRecentIp)
		}
		time.Sleep(3600 * time.Second)
	}
}
