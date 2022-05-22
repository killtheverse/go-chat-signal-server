package config

import "os"

type Config struct {
    ServerAddress   string
    MongoURI        string
    MongoDBName     string
    SecretKey       string
    RedisURL        string
}

// ReadConfig returns a new Config instance
func ReadConfig() *Config {
    return &Config{
        ServerAddress: ":" + os.Getenv("PORT"),
        MongoURI: os.Getenv("MONGO_URI"),
        MongoDBName: os.Getenv("MONGO_DB_NAME"),
        SecretKey: os.Getenv("SECRET_KEY"),
        RedisURL: os.Getenv("REDIS_URL"),
    }
}
