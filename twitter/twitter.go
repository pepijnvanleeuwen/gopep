// Copyright (c) 2016 Pepijn van Leeuwen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package twitter

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/pepijnvanleeuwen/gopep/configuration"
)

var api *anaconda.TwitterApi
var config payload

type payload struct {
	consumerKey       string
	consumerSecret    string
	accessToken       string
	accessTokenSecret string
}

func init() {
	payload, err := configuration.GetPayload(configuration.Twitter)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(payload.Data), &config); err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(config.consumerKey)
	anaconda.SetConsumerSecret(config.consumerSecret)
	api = anaconda.NewTwitterApi(config.accessToken, config.accessTokenSecret)
}

// Load the specified action by passing the provided value.
func Load(action, value string) error {
	switch action {
	case "like":
	case "l":
		return Like(value)
	}

	return fmt.Errorf("Action '%s' is not supported.", action)
}

// Like the provided tweet ID.
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
