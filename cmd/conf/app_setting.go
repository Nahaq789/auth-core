package conf

type AppSetting struct {
	Server ServerSetting
	Aws    AwsSetting
}

type AwsSetting struct {
	Region string `env:"REGION"`
}

type ServerSetting struct {
	Port string `env:"PORT"`
}
