package methods

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Delete(url string) {

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		log.Println(err)
	}

	

	fmt.Println("Response Body:\n", resp)
	fmt.Println(" Status Code: ", resp.StatusCode())

}