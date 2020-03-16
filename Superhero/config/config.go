package config

type Config struct {
	DB *DBConfig
	DB1 *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "jamespdh",
			Name:     "superheromovie",
			Charset:  "utf8",
		},
		DB1: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "jamespdh",
			Name:     "topcollection",
			Charset:  "utf8",
		},
	}
}


