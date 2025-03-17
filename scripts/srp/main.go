package main

import (
	"fmt"
	"os"

	cognitosrp "github.com/alexrudd/cognito-srp/v4"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Warning: failed to load .env file")
	}

	// 環境変数から認証情報を取得
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	poolId := os.Getenv("POOLID")
	clientId := os.Getenv("CLIENT_ID")

	// CLIENT_SECRETが設定されているか確認
	var clientSecretPtr *string
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret != "" {
		clientSecretPtr = &clientSecret
	} else {
		clientSecretPtr = nil
	}

	// SRPクライアントを初期化
	csrp, err := cognitosrp.NewCognitoSRP(username, password, poolId, clientId, clientSecretPtr)
	if err != nil {
		fmt.Printf("Error initializing CognitoSRP: %v\n", err)
		return
	}

	// 認証パラメータを取得
	authParams := csrp.GetAuthParams()

	// 見やすく出力
	fmt.Println("=== AWS Cognito SRP Parameters ===")
	fmt.Printf("USERNAME: %s\n", authParams["USERNAME"])
	fmt.Printf("SRP_A: %s\n", authParams["SRP_A"])

	// SRP_Aを単独で取得しやすいようにする
	fmt.Println("\n=== SRP_A Value Only (For Copy/Paste) ===")
	fmt.Println(authParams["SRP_A"])

	return
}
