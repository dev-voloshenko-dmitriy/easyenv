package easyenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvRequired(t *testing.T) {
	envDefault := EnvRequired{}

	assert.NotNil(t, envDefault.GetKeyHelper())
}

func TestExecute(t *testing.T) {
	envDefault := EnvRequired{}
	incoming_value := "test"
	reqDefault, err := envDefault.Execute("true", incoming_value)
	assert.NoError(t, err)
	assert.Equal(t, incoming_value, reqDefault)

	reqDefault, err = envDefault.Execute("false", "")
	assert.NoError(t, err)
	assert.Equal(t, "", reqDefault)

	reqDefault, err = envDefault.Execute("true", "")
	assert.Error(t, err)
}
