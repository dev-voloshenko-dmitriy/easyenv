package easyenv

import (
	"errors"
	"fmt"
)

type EnvRequired struct{}

func (e *EnvRequired) GetKeyHelper() string {
	return "env-required"
}

func (e *EnvRequired) Execute(value string, incoming_value string) (string, error) {

	if value == "true" && incoming_value == "" {
		return "", errors.New(fmt.Sprintf("Environment variable %s is required", e.GetKeyHelper()))
	}

	return incoming_value, nil
}
