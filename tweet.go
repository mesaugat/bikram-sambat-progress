package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dghubble/oauth1"
)

// TwitterCredentials stores all access/consumer tokens and secret keys needed
// for authentication against the twitter REST API.
type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// Tweet uses twitter credentials to post a tweet using Twitter v2 API.
// https://developer.twitter.com/en/docs/twitter-api/tweets/manage-tweets/api-reference/post-tweets
func Tweet(text string) {
	creds := TwitterCredentials{
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
	}

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/2/tweets"
	reqBody := []byte(fmt.Sprintf(`{"text": "%s"}`, text))
	resp, err := httpClient.Post(path, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Printf("Error making http request to post a tweet %s\n", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	log.Printf("Posted a Tweet")
	log.Printf("HTTP Response: %s", string(body))
}
