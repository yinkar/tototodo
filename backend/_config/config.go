package config

type Config struct {
	DB  *DBConfig
	Srv *ServerConfig
	Api *ApiConfig
}

type DBConfig struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Charset  string
}

type ServerConfig struct {
	Port int
}

type ApiConfig struct {
	Version string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Driver:   "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "root",
			Database: "tototodo",
			Charset:  "utf8",
		},
		Srv: &ServerConfig{
			Port: 8000,
		},
		Api: &ApiConfig{
			Version: "v0.1.0",
		},
	}
}
