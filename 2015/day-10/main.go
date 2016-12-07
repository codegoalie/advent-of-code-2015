package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Message is the heding text and attachments for the Slack message
type Message struct {
	Text string `json:"text"`
}

func main() {
	seen := "1113122113"
	for i := 0; i < 50; i++ {
		seen = lookAndSee(seen)
		fmt.Printf("i = %+v\n", i)
		postToSlack(Message{Text: fmt.Sprintf("i = %+v\n", i)})
	}
	fmt.Printf("len(seen) = %+v\n", len(seen))
	postToSlack(Message{Text: fmt.Sprintf("len(seen) = %+v\n", len(seen))})
}

func lookAndSee(in string) string {
	chars := strings.Split(in, "")
	current := chars[0]
	count := 0
	saw := ""

	for _, char := range chars {
		if current != char {
			saw += strconv.Itoa(count) + current
			current = char
			count = 1
		} else {
			count++
		}
	}
	saw += strconv.Itoa(count) + current
	return saw
}

func postToSlack(message Message) {
	b, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(b))

	resp, err := http.PostForm(os.Getenv("SLACK_URL"),
		url.Values{"payload": {string(b)}})

	if err != nil {
		fmt.Println("Error posting to Slack", err)
	} else {
		rawBody, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			fmt.Println("failed to read Slack response", err)
		} else {
			fmt.Println("Slack response", string(rawBody))
		}
	}
}
