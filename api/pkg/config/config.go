package config

import (
	"fmt"
	"os"
)

type env struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	ENVIRONMENT       string
}

var Env env

func init() {
	var err error
	var envErrs []error

	if Env.POSTGRES_USER, err = getEnv("POSTGRES_USER"); err != nil {
		envErrs = append(envErrs, err)
	}
	if Env.POSTGRES_PASSWORD, err = getEnv("POSTGRES_PASSWORD"); err != nil {
		envErrs = append(envErrs, err)
	}
	if Env.POSTGRES_DB, err = getEnv("POSTGRES_DB"); err != nil {
		envErrs = append(envErrs, err)
	}
	if Env.POSTGRES_HOST, err = getEnv("POSTGRES_HOST"); err != nil {
		envErrs = append(envErrs, err)
	}
	if Env.ENVIRONMENT, err = getEnv("ENVIRONMENT"); err != nil {
		envErrs = append(envErrs, err)
	}

	if len(envErrs) != 0 {
		var errMsg string
		for _, envErr := range envErrs {
			errMsg += fmt.Sprintf("%s\n", envErr.Error())
		}
		panic(errMsg)
	}
}

func getEnv(key string) (string, error) {
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("getEnv: %s not found", key)
	}
	return e, nil
}
