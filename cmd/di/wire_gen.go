// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/auth-core/internal/application"
	"github.com/auth-core/internal/application/service"
	repository2 "github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/infrastructure/repository"
	"github.com/auth-core/internal/presentation/controller"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Initialize(client *dynamodb.Client) *ControllerSet {
	string2 := _wireStringValue
	userRepositoryImpl := repository.NewUserRepositoryImpl(client, string2)
	userServiceImpl := services.NewUserService(userRepositoryImpl)
	authController := controller.NewAuthController(userServiceImpl)
	diControllerSet := &ControllerSet{
		AuthController: authController,
	}
	return diControllerSet
}

var (
	_wireStringValue = "hoge"
)

// wire.go:

var awsSet = wire.NewSet(dynamodb.New)

var repositorySet = wire.NewSet(repository.NewUserRepositoryImpl, wire.Bind(new(repository2.UserRepository), new(*repository.UserRepositoryImpl)), wire.Value("hoge"))

var serviceSet = wire.NewSet(services.NewUserService, wire.Bind(new(application.UserService), new(*services.UserServiceImpl)))

var controllerSet = wire.NewSet(controller.NewAuthController)

type ControllerSet struct {
	AuthController *controller.AuthController
}
