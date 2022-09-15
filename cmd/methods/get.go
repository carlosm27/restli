package methods

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func Get(url string) {
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Response Body:\n", resp)
	fmt.Println(" Status Code: ", resp.StatusCode())

}
