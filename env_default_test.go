package easyenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetEnv(t *testing.T) {
	envDefault := EnvDefault{}

	if envDefault.GetKeyHelper() == "" {
		t.Error("KeyHelper is empty")
	}
}

func Test_Execute(t *testing.T) {
	envDefault := EnvDefault{}
	defaultVal := "default"

	executeResult, err := envDefault.Execute(defaultVal, "")

	assert.NoError(t, err)
	assert.Equal(t, defaultVal, executeResult)

	executeResultNormal, err := envDefault.Execute(defaultVal, "test")
	assert.NoError(t, err)
	assert.Equal(t, "test", executeResultNormal)
}
