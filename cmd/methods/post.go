package methods

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Post(url string, body string) {

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(body)

	fmt.Println("Response Body:\n", resp)
	fmt.Println(" Status Code: ", resp.StatusCode())

}
