package env

import (
	"os"
)

func MustEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing env var: " + key)
	}
	return val
}