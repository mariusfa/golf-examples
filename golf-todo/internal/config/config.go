package config

import "github.com/mariusfa/golf/database"

type Config struct {
	Port           string
	DbHost         string
	DbPort         string
	DbUser         string
	DbPassword     string
	DbName         string
	DbAppUser      string
	DbAppPassword  string
	DbRunBaseLine  string `required:"false"`
	DbStartupLocal string `required:"false"`
}

func (c *Config) ToDbConfig() *database.DbConfig {
	return database.NewDbConfig(
		c.DbHost,
		c.DbName,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbAppUser,
		c.DbAppPassword,
		c.DbRunBaseLine,
		c.DbStartupLocal,
	)
}
