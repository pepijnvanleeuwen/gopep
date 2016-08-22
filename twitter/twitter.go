package twitter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pepijnvanleeuwen/gopep/configuration"
)

var api *anaconda.TwitterApi
var config payload
var initialized bool

type payload struct {
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
}

func initialize() {
	log.Println("Initializing 'twitter' module")

	payload, err := configuration.GetPayload(configuration.Twitter)
	if err != nil {
		log.Fatalln(err)
	}

	data, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalln(err)
	}

	anaconda.SetConsumerKey(config.consumerKey)
	anaconda.SetConsumerSecret(config.consumerSecret)
	api = anaconda.NewTwitterApi(config.accessToken, config.accessTokenSecret)

	initialized = true
}

// Load loads the specified action by passing the provided value.
func Load(action, value string) error {
	if !initialized {
		initialize()
	}

	switch action {
	case "like", "l":
		return Like(value)
	case "":
		return fmt.Errorf("Please specify the action.")
	}

	return fmt.Errorf("Action '%s' is not supported.", action)
}

// Like likes the provided tweet ID.
func Like(v string) error {
	if v == "" {
		return fmt.Errorf("Cannot like tweet - no identifier provided.")
	}

	id, err := strconv.ParseInt(v, 16, 64)
	if err != nil {
		return err
	}

	tweet, err := api.Favorite(id)
	if err != nil {
		return err
	}

	fmt.Printf("Liked tweet submitted by %s at %s:", tweet.User.Name, tweet.CreatedAt)
	fmt.Println(tweet.Text)

	return nil
}
