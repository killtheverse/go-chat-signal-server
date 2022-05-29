package config


type MongodbConfig struct {
    DatabaseName            string      `mapstructure:"db_name"`
    TimeOut                 int         `mapstructure:"timeout"`
    DialTimeOut             int64       `mapstructure:"dial_timeout"`
    URI                     string      `mapstructure:"uri"`
}
