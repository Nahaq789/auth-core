//go:build wireinject
// +build wireinject

package di

import (
	"github.com/auth-core/cmd/conf"
	"github.com/auth-core/internal/application"
	services "github.com/auth-core/internal/application/service"
	domainRepos "github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/infrastructure/repository"
	"github.com/auth-core/internal/presentation/controller"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

func ProvideUserRepository(client *dynamodb.Client, aws *conf.AwsSetting) *repository.UserRepositoryImpl {
	repository := repository.NewUserRepositoryImpl(client, aws.UserTable)
	return repository
}

func ProvideCognitoRepository(client *cognitoidentityprovider.Client, clientId string) *repository.CognitoRepositoryImpl {
	repository := repository.NewCognitoRepository(client, clientId)
	return repository
}

var awsSet = wire.NewSet(dynamodb.New)

var repositorySet = wire.NewSet(
	ProvideUserRepository,
	wire.Bind(new(domainRepos.UserRepository), new(*repository.UserRepositoryImpl)),
)
var CognitoSet = wire.NewSet(
	ProvideCognitoRepository,
	wire.Bind(new(domainRepos.CognitoRepository), new(*repository.CognitoRepositoryImpl)),
)

var serviceSet = wire.NewSet(
	services.NewUserService,
	wire.Bind(new(application.UserService), new(*services.UserServiceImpl)),
)
var controllerSet = wire.NewSet(controller.NewAuthController)

type ControllerSet struct {
	AuthController *controller.AuthController
}

func Initialize(client *dynamodb.Client, aws *conf.AwsSetting) *ControllerSet {
	wire.Build(
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil
}
