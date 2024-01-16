package getenv

import "os"

func GetEnv(name string) string {
	result := os.Getenv(name)
	return result
}
