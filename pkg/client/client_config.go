package client

import (
	"fmt"
	"os"

	azaci "github.com/Azure/azure-sdk-for-go/services/containerinstance/mgmt/2021-10-01/containerinstance"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
)

type Config struct {
	ClientID             string
	ClientSecret         string
	UserIdentityClientId string
	TenantID             string
	SubscriptionID       string
	lClient              azaci.LocationClient
}

func (config *Config) InitClient() {
	if config.SubscriptionID == "" {
		fmt.Println("subscriptionID cannot be empty")
		os.Exit(1)
	}

	config.lClient = azaci.NewLocationClientWithBaseURI("https://management.azure.com/", config.SubscriptionID)
	auth, err := config.getAuthorizer()
	if err != nil {
		fmt.Println("an error has occurred while getting the authorizer ", err)
		os.Exit(1)
	}
	config.lClient.Authorizer = auth
	fmt.Println("client is ready")

}

// getAuthorizer return autorest authorizer.
func (config *Config) getAuthorizer() (autorest.Authorizer, error) {
	fmt.Println("getting authorizer")

	var auth autorest.Authorizer
	var err error
	resource := "https://management.azure.com/"

	var token *adal.ServicePrincipalToken
	isUserIdentity := len(config.ClientID) == 0

	if isUserIdentity {
		fmt.Println("using managed identity to get token")

		if config.UserIdentityClientId == "" {
			fmt.Println("userIdentityClientId cannot be empty")
			os.Exit(1)
		}

		token, err = adal.NewServicePrincipalTokenFromManagedIdentity(resource, &adal.ManagedIdentityOptions{ClientID: config.UserIdentityClientId})
		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("using service principal to get token")

		if config.ClientID == "" || config.ClientSecret == "" {
			fmt.Println("clientID or clientSecret cannot be empty")
			os.Exit(1)
		}

		if config.TenantID == "" {
			fmt.Println("tenantID cannot be empty")
			os.Exit(1)
		}

		oauthConfig, err := adal.NewOAuthConfig(
			"https://login.microsoftonline.com/", config.TenantID)
		if err != nil {
			return nil, err
		}

		token, err = adal.NewServicePrincipalToken(
			*oauthConfig, config.ClientID, config.ClientSecret, resource)
		if err != nil {
			return nil, err
		}
	}

	auth = autorest.NewBearerAuthorizer(token)
	return auth, err
}
