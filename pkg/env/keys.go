package env

import (
	"fmt"
	"os"
)

type Key string

type Env string

func (k Key) GetEnv() (Env, error) {
	env := os.Getenv(string(k))
	if env == "" {
		return "", fmt.Errorf("env %s is not set", k)
	}

	return Env(env), nil
}

func (e Env) String() string {
	return string(e)
}
