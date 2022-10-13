package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"listCap/pkg/client"
)

func enterInput(param string) string {
	fmt.Printf("Please enter the %s: ", param)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading ", param, ". Please try again", err)
		return ""
	}
	result := strings.TrimSuffix(input, "\n")
	fmt.Println(result)
	return result
}

func main() {

	config := client.Config{}
	config.SubscriptionID = enterInput("SubscriptionID")
	config.TenantID = enterInput("TenantID")
	region := enterInput("region")

	fmt.Println("** Next, either enter the ClientID & ClientSecret OR UserIdentityClientId **")
	config.ClientID = enterInput("ClientID")
	config.ClientSecret = enterInput("ClientSecret")
	config.UserIdentityClientId = enterInput("UserIdentityClientId")

	config.InitClient()

	if region == "" {
		fmt.Println("region input is mandatory")
		os.Exit(1)
	}
	result, err := config.ListCapabilities(context.TODO(), region)
	if err != nil {
		fmt.Println("an error has occurred while calling ListCapabilities", err)
		os.Exit(1)
	}
	fmt.Println("result returned successfully", result)
}
