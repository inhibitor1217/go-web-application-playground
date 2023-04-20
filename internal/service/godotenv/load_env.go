package godotenv

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	BASE_ENV_FILE = "config/.env"
	ENV_FILE      = "config/.env.%s"
)

func LoadEnv() {
	godotenv.Load(
		fmt.Sprintf(ENV_FILE, os.Getenv("STAGE")),
		BASE_ENV_FILE,
	)
}
