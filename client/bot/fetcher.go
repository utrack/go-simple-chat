package clientBot

import (
	"github.com/pquerna/ffjson/ffjson"
	"io/ioutil"
	"net/http"
	"time"
)

type chuckApiModel struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}

const chuckUrl = `http://api.icndb.com/jokes/random`

func getQuote() (string, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := netClient.Get(chuckUrl)
	if err != nil {
		return ``, err
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ``, err
	}

	var joke chuckApiModel
	err = ffjson.Unmarshal(buf, &joke)
	if err != nil {
		return ``, err
	}

	return joke.Value.Joke, nil
}
