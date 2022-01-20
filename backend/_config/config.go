package config

type Config struct {
	DB  *DBConfig
	Srv *ServerConfig
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
	}
}
