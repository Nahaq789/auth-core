// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/auth-core/cmd/conf"
	"github.com/auth-core/internal/application"
	"github.com/auth-core/internal/application/service"
	repository2 "github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/infrastructure/repository"
	"github.com/auth-core/internal/presentation/controller"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
	"log/slog"
)

// Injectors from wire.go:

func Initialize(logger *slog.Logger, dynamodb2 *dynamodb.Client, cognito *cognitoidentityprovider.Client, aws *conf.AwsSetting) *ControllerSet {
	userRepositoryImpl := ProvideUserRepository(logger, dynamodb2, aws)
	cognitoRepositoryImpl := ProvideCognitoRepository(cognito, aws)
	userServiceImpl := services.NewUserService(logger, userRepositoryImpl, cognitoRepositoryImpl)
	authController := controller.NewAuthController(userServiceImpl)
	diControllerSet := &ControllerSet{
		AuthController: authController,
	}
	return diControllerSet
}

// wire.go:

func ProvideUserRepository(logger *slog.Logger, client *dynamodb.Client, aws *conf.AwsSetting) *repository.UserRepositoryImpl {
	repository2 := repository.NewUserRepositoryImpl(client, aws.UserTable)
	return repository2
}

func ProvideCognitoRepository(client *cognitoidentityprovider.Client, aws *conf.AwsSetting) *repository.CognitoRepositoryImpl {
	repository2 := repository.NewCognitoRepository(client, aws.CognitoClientId)
	return repository2
}

var awsSet = wire.NewSet(dynamodb.New)

var repositorySet = wire.NewSet(
	ProvideUserRepository, wire.Bind(new(repository2.UserRepository), new(*repository.UserRepositoryImpl)),
)

var CognitoSet = wire.NewSet(
	ProvideCognitoRepository, wire.Bind(new(repository2.CognitoRepository), new(*repository.CognitoRepositoryImpl)),
)

var serviceSet = wire.NewSet(services.NewUserService, wire.Bind(new(application.UserService), new(*services.UserServiceImpl)))

var controllerSet = wire.NewSet(controller.NewAuthController)

type ControllerSet struct {
	AuthController *controller.AuthController
}
