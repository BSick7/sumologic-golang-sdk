package main

import (
	"fmt"
	"github.com/BSick7/sumologic-sdk-go/api"
)

func main() {
	session := api.DefaultSession()
	client := api.NewClient(session)
	collectors, err := client.Collectors().List(0, 5)
	fmt.Println(collectors, err)
}
