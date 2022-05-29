package config


type ServerConfig struct {
    ReadTimeOut         int       `mapstructure: "readTimeOut"`
    WriteTimeOut        int       `mapstructure: "writeTimeOut"`
    IdleTimeOut         int       `mapstructure: "idleTimeOut"`
}

