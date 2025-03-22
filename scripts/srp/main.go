package main

import (
	"context"
	"fmt"
	"os"
	"time"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
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

	ctx := context.Background()
	sdkConfig, _ := config.LoadDefaultConfig(ctx)
	client := cognitoidentityprovider.NewFromConfig(sdkConfig)

	resp, err := client.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowTypeUserSrpAuth,
		ClientId:       aws.String(csrp.GetClientId()),
		AuthParameters: authParams,
	})
	if err != nil {
		panic(err)
	}

	if resp.ChallengeName == types.ChallengeNameTypePasswordVerifier {
		challengeResponse, _ := csrp.PasswordVerifierChallenge(resp.ChallengeParameters, time.Now())

		resp, err := client.RespondToAuthChallenge(ctx, &cognitoidentityprovider.RespondToAuthChallengeInput{
			ChallengeName:      types.ChallengeNameTypePasswordVerifier,
			ChallengeResponses: challengeResponse,
			ClientId:           aws.String(csrp.GetClientId()),
		})

		if err != nil {
			panic(err)
		}

		fmt.Println()
		fmt.Println("=== Tokens ===")
		fmt.Printf("Access Token: %s\n", *resp.AuthenticationResult.AccessToken)
		fmt.Printf("ID Token: %s\n", *resp.AuthenticationResult.IdToken)
		fmt.Printf("Refresh Token: %s\n", *resp.AuthenticationResult.RefreshToken)
	}
	return
}
