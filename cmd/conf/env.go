package conf

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Load() (*AppSetting, error) {
	conf := &AppSetting{}
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		if err := env.Parse(conf); err != nil {
			log.Fatal(err)
		}
	} else {
		_, b, _, _ := runtime.Caller(0)
		rootPath := filepath.Join(filepath.Dir(b), "../../")
		if err := godotenv.Load(filepath.Join(rootPath, ".env")); err != nil {
			log.Fatal(err)
		}

		if err := env.Parse(conf); err != nil {
			log.Fatal(err)
		}
	}

	return conf, nil
}
