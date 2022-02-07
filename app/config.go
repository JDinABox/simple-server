package app

type Config struct {
	Port int
}

func (c *Config) Default() {
	if c.Port == 0 {
		c.Port = 8080
	}
}
