package conf

type AppSetting struct {
	Server ServerSetting
	Aws    AwsSetting
}

type AwsSetting struct {
	Region    string `env:"REGION"`
	UserTable string `env:"USER_TABLE"`
}

type ServerSetting struct {
	Port  string `env:"PORT"`
	Level string `enc:"LEVEL"`
}
