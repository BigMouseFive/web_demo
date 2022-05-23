package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	Config *viper.Viper
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

// Create 创建配置
func Create() {
	Config = viper.New()
	configFile := os.Getenv("CONFIG_FILE")
	if len(configFile) == 0 {
		configFile = "webapi"
	}
	setDefault()
	Config.SetConfigName("webapi")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("/config")
	Config.AddConfigPath("./src")
	Config.AddConfigPath("./")
	err := Config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	Config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	Config.WatchConfig()
}

func setDefault() {
	Config.SetDefault("host", "")
	Config.SetDefault("redis.address", "127.0.0.1")
	Config.SetDefault("redis.password", "Arrive@2016reDis#")
	Config.SetDefault("mongo.address", "mongodb://127.0.0.1:27017")
	Config.SetDefault("mongo.username", "demo")
	Config.SetDefault("mongo.password", "arrive@2016")
	Config.SetDefault("mongo.db_name", "demo")
	Config.SetDefault("sync.dir", "www.arrive.ai:/nfs/demo/sync/")
	Config.SetDefault("sync.password", "aRRive@2021A")
	Config.SetDefault("mqtt.address", "192.168.2.102:1883")
	Config.SetDefault("static.relativePath", "static")
	Config.SetDefault("static.root", "/data/static")
	Config.SetDefault("static.dataManagerId", "SyncData001")
}
