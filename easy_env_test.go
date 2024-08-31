package easyenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DefaultTest struct {
	A string `env:"env" env-default:"local"`
	B int    `env:"port" env-required:"true"`
}

type FailInt64 struct {
	B string `env:"env" env-default:"local"`
	A int64  `env:"port"`
}

type FloatTest struct {
	B string  `env:"env" env-default:"local"`
	A float64 `env:"port"`
}

func TestNewDecoder(t *testing.T) {
	decoder, err := NewDecoder("./.env")
	assert.NoError(t, err)
	assert.IsType(t, &Env{}, decoder)

	decoder, err = NewDecoder("./not_exist_file.env")
	assert.Error(t, err)
}

func TestLoad(t *testing.T) {
	decoder, err := NewDecoder("./.env")
	assert.NoError(t, err)

	defaultTest := DefaultTest{}

	err = decoder.Load(&defaultTest)
	assert.NoError(t, err)

	failInt64 := FailInt64{}

	err = decoder.Load(&failInt64)
	assert.Error(t, err)

	err = decoder.Load(failInt64)
	assert.Error(t, err)

	err = decoder.Load(&FloatTest{})
	assert.NoError(t, err)

	decoder, err = NewDecoder("./.env_empty")
	assert.NoError(t, err)

	err = decoder.Load(&defaultTest)
	assert.Error(t, err)

	decoder, err = NewDecoder("./test.json")
	assert.Error(t, err)
}
