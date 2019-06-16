package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Gets the latest Tweet and returns its tweet.Text
func getTweet() string {
	token := os.Getenv("TWITTER_TOKEN")
	tokenSecret := os.Getenv("TWITTER_TOKEN_SECRET")
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	twitterScreenName := os.Getenv("TWITTER_SCREEN_NAME")
	api := anaconda.NewTwitterApiWithCredentials(token, tokenSecret, consumerKey, consumerSecret)
	twitter := url.Values{}
	twitter.Set("screen_name", twitterScreenName)
	twitter.Add("count", "1")
	tweets, _ := api.GetUserTimeline(twitter)
	return tweets[0].Text
}

// Main
func main() {
	fmt.Println(getTweet())
	return
}
