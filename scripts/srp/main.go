package main

import (
	"fmt"
	"os"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Warning: failed to load .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	poolId := os.Getenv("POOLID")
	clientId := os.Getenv("CLIENT_ID")

	var clientSecretPtr *string
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret != "" {
		clientSecretPtr = &clientSecret
	} else {
		clientSecretPtr = nil
	}

	csrp, err := cognitosrp.NewCognitoSRP(username, password, poolId, clientId, clientSecretPtr)
	if err != nil {
		fmt.Printf("Error initializing CognitoSRP: %v\n", err)
		return
	}

	authParams := csrp.GetAuthParams()

	fmt.Println("=== AWS Cognito SRP Parameters ===")
	fmt.Printf("AUTH_PARAMS: %v\n", authParams)
	fmt.Printf("USERNAME: %s\n", authParams["USERNAME"])
	fmt.Printf("SRP_A: %s\n", authParams["SRP_A"])

	fmt.Println("\n=== SRP_A Value Only (For Copy/Paste) ===")
	fmt.Println(authParams["SRP_A"])

	return
}
