package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)


	
type Status struct {
	YsStatusMessageKey     string   `json:"ysStatusMessageKey"`
}


func sendWebhook () {
	url := "https://discord.com/api/webhooks/846657090192801813/qD7LZbFgWESQvPatifyjenECqV1Wpc26zLroSK25XsX7siqq6FXR7yam0u7SsI_NXK7G"

	var jsonStr = []byte(`{"content":null,"embeds":[{"title":"YS Sale Status","description":"Sale has started 🎉 @everyone","color":5814783}]}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}



func main () {
	saleStarted := false

	for {
		client := &http.Client{}

		request, requestError := http.NewRequest("GET", "http://localhost:3333/hpl/content/yeezy-supply/config/US/waitingRoomConfig.json", nil)

		if requestError != nil {
			fmt.Println(requestError)
			return
		}

		request.Close = true

		headers := map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
			"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
			"cache-control": "no-store,no-cache,must-revalidate,proxy-revalidate,max-age=0",
			"pragma": "no-cache",
		}

		for key, value := range headers {
			request.Header.Set(key, value)
			
		}

		response, responseError := client.Do(request)
		if responseError != nil {
			fmt.Println(responseError)
			return
		}

		var status Status

		body, bodyErr := ioutil.ReadAll(response.Body)

		if bodyErr != nil {
			fmt.Println(bodyErr)
			return
		}

		response.Body.Close()

		err := json.Unmarshal(body, &status)

		if err != nil {
			fmt.Println(err)
			return
		}

		if status.YsStatusMessageKey == "sale_started" {
			if !saleStarted {
				saleStarted = true
				sendWebhook()
			}
		} else {
			saleStarted = false
		}

		time.Sleep(3 * time.Second)
	}
}