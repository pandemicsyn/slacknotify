package slacknotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//SlackNotify sends messages to slack
type SlackNotify struct {
	URL string
}

type slackMsg struct {
	Text string `json:"text"`
}

//Send a message to slack, returns an error on failure
func (s *SlackNotify) Send(v ...interface{}) error {
	payload, err := json.Marshal(slackMsg{Text: fmt.Sprint(v)})
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", s.URL, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("Error posting to slack")
	}
	return nil
}

//Println sends a message to slack, self logs an error on internal failure
func (s *SlackNotify) Println(v ...interface{}) {
	payload, err := json.Marshal(slackMsg{Text: fmt.Sprint(v)})
	if err != nil {
		log.Println(err.Error())
	}
	req, _ := http.NewRequest("POST", s.URL, bytes.NewBuffer(payload))
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("Error posting to slack", res.StatusCode)
	}
	if err != nil {
		log.Println(err.Error())
	}
}
