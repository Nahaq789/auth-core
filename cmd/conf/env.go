package conf

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadAppSetting() (*AppSetting, error) {
	server := &ServerSetting{}
	aws := &AwsSetting{}
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		if err := env.Parse(server); err != nil {
			log.Fatal(err)
		}
		if err := env.Parse(aws); err != nil {
			log.Fatal(err)
		}
	} else {
		_, b, _, _ := runtime.Caller(0)
		rootPath := filepath.Join(filepath.Dir(b), "../../")
		if err := godotenv.Load(filepath.Join(rootPath, ".env")); err != nil {
			log.Fatal(err)
		}
		if err := env.Parse(server); err != nil {
			log.Fatal(err)
		}
		if err := env.Parse(aws); err != nil {
			log.Fatal(err)
		}
	}

	app := AppSetting{
		Aws:    *aws,
		Server: *server,
	}
	return &app, nil
}
