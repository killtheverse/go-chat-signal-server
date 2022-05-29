package injectors

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/killtheverse/go-chat-signal-server/config"
	"github.com/spf13/viper"
)

func ProvideConfig(config_dir string) (*config.Config, error) {
    files, err := ioutil.ReadDir(config_dir)
    if err != nil {
        panic(fmt.Errorf("Can't open config directory\n"))
    }
    if len(files) == 0 {
        panic(fmt.Errorf("Empty config directory\n"))
    }
    file := files[0]
    viper.SetConfigFile(fmt.Sprintf("%s/%s", config_dir, file.Name()))
    viper.SetConfigType("yaml")
    err = viper.ReadInConfig()
    if err != nil {
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    
    for _, file := range files[1:] {
        f, err := os.Open(fmt.Sprintf("%s/%sv", config_dir, file.Name()))
        if err != nil {
            panic(fmt.Errorf("Fatal error reading config file: %s\n", err))
        }
        defer f.Close()
        err = viper.MergeConfig(f)
        if err != nil {
            panic(fmt.Errorf("Fatal error merging config file: %s\n", err))
        }
    }

    var conf config.Config
    err = viper.Unmarshal(&conf)
    if err != nil {
        return nil, fmt.Errorf("Fatal error in unmarshalling config file: %s\n", err)
    }
    return &conf, nil
} 
