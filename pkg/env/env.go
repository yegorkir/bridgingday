package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	if err := godotenv.Load(path); err != nil {
		return fmt.Errorf("file: %s, err: %v", path, err)
	}

	return nil
}
