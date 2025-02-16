package main

import (
	"context"
	"fmt"

	"github.com/auth-core/cmd/conf"
	"github.com/auth-core/cmd/di"
	"github.com/auth-core/pkg/db"
	"github.com/gin-gonic/gin"
)

func Routing(r gin.IRouter, settng *conf.AppSetting) error {
	v1 := r.Group("/api/v1")
	client, err := db.NewDynamoDbClient(context.Background(), settng.Aws.Region)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	cr := di.Initialize(client)
	{
		v1.POST("/signup", cr.AuthController.Signup)
	}

	return nil
}
