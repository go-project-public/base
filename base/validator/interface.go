package validator

func Get() *Config {
	c := &Config{}
	if err := c.Init(); err != nil {
		panic(err)
	}
	return c
}

func Validate(target any) error {
	return Get().Validate(target)
}
