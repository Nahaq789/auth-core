//go:build wireinject
// +build wireinject

package di

import (
	"github.com/auth-core/internal/application"
	services "github.com/auth-core/internal/application/service"
	domainRepos "github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/infrastructure/repository"
	"github.com/auth-core/internal/presentation/controller"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/wire"
)

var awsSet = wire.NewSet(dynamodb.New)
var repositorySet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(domainRepos.UserRepository), new(*repository.UserRepositoryImpl)),
	wire.Value("hoge"),
)
var serviceSet = wire.NewSet(
	services.NewUserService,
	wire.Bind(new(application.UserService), new(*services.UserServiceImpl)),
)
var controllerSet = wire.NewSet(controller.NewAuthController)

type ControllerSet struct {
	AuthController *controller.AuthController
}

func Initialize(client *dynamodb.Client) *ControllerSet {
	wire.Build(
		repositorySet,
		serviceSet,
		controllerSet,
		wire.Struct(new(ControllerSet), "*"),
	)
	return nil
}
