package godotenv

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const (
	BASE_ENV_FILE = "config/.env"
	ENV_FILE      = "config/.env.%s"
)

var Option = fx.Options(
	fx.Invoke(func() {
		godotenv.Load(
			fmt.Sprintf(ENV_FILE, os.Getenv("STAGE")),
			BASE_ENV_FILE,
		)
	}),
)
