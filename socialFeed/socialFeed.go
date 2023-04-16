package socialFeed

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FacebookPosts() []byte {

	url := "https://graph.facebook.com/v16.0/me/posts?access_token="
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	return body
}
