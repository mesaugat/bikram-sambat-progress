package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// TwitterCredentials stores all access/consumer tokens
// and secret keys needed for authentication against
// the twitter REST API.
type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// getTwitterClient uses twitter credentials to give back the twitter client
// https://tutorialedge.net/golang/writing-a-twitter-bot-golang/
func getTwitterClient() *twitter.Client {
	creds := TwitterCredentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// Retrieve the user and verify if the credentials are valid
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		log.Println(err)
		panic("Cannot verify Twitter credentials")
	}

	log.Printf("Twitter User \n%+v\n", user)

	return client
}

// Tweet tweets a message
func Tweet(message string) (*twitter.Tweet, error) {
	twitter := getTwitterClient()

	tweet, resp, err := twitter.Statuses.Update(message, nil)

	if err != nil {
		log.Println("Cannot update Twitter status")
		log.Println(err)

		return nil, err
	}

	log.Printf("Response \n%+v\n", resp)
	log.Printf("Tweet \n%+v\n", tweet)

	return tweet, nil
}
