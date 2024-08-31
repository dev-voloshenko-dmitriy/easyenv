package easyenv

type EnvDefault struct{}

func (e *EnvDefault) GetKeyHelper() string {
	return "env-default"
}

func (e *EnvDefault) Execute(value string, incoming_value string) (string, error) {
	if incoming_value == "" {
		return value, nil
	}

	return incoming_value, nil
}
