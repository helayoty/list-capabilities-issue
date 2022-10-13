# listCapabilities Issue

This command utility to help us reproduce the listCapabilities issue when using Managed Identity.

1. Clone the repo

```shell
$ git clone https://github.com/helayoty/list-capabilities-issue.git

$ cd listCap
```

2. Run the command utility

```shell
$ go run main.go
```

3. Enter the following inputs

> **Note** `SubscriptionID`, `TenantID`, `region` are mandatory inputs, while either (`ClientID` & `ClientSecret`) `UserIdentityClientId` is optional 

```shell
Please enter the SubscriptionID: #####
#####
Please enter the TenantID: #####
#####
Please enter the region: <enter the region>

** Next, either enter the ClientID & ClientSecret OR UserIdentityClientId **
Please enter the ClientID:

Please enter the ClientSecret: 

Please enter the UserIdentityClientId: #####
#####
```

4. In case of using the service principal (`ClientID` & `ClientSecret`), you can see the result printed out with a successful message.

5. In case of Managed Identity (`UserIdentityClientId`), the program will hang.