package config


type Config struct {
    Server          ServerConfig            `mapstructure:"server"`
    MongoDB         MongodbConfig           `mapstructure:"mongodb"`
}
