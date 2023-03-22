package aigenRest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func TwitterHome() []byte {

	url := "https://api.twitter.com/2/users/999662318186303488/timelines/reverse_chronological?max_results=100"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []byte("")
	}
	req.Header.Add("oauth_consumer_key", "D07M43AMAztnVxNfqqXQuXFio")
	req.Header.Add("oauth_signature", "Uoo1D1nNOsIzjB1E7FIGHOfHh3SrYXJ9zd7r9jQHEk36VuIc")
	req.Header.Add("Authorization", "OAuth oauth_consumer_key=\"D07M43AMAztnVxNfqqXQuXFio\",oauth_token=\"999662318186303488-65J434J2dxSXr3XP4XuktQ4YcETGHmj\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1679439270\",oauth_nonce=\"xgzZ5Tu5k2C\",oauth_version=\"1.0\",oauth_signature=\"r5BF8xM2HHlKAUmABqN3tYmfkuM%3D\"")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte("")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte("")
	}
	fmt.Println(string(body))
	return body
}

func SendTweet(tweet string) {

	url := "https://api.twitter.com/2/tweets"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
    "text": "%s"
}`, tweet))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
