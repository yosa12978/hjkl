package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

var (
	cfg  Config
	once sync.Once
)

type Config struct {
	App struct {
		Host         string `yaml:"host" envconfig:"HJKL_HOST" json:"host"`
		Port         int    `yaml:"port" envconfig:"HJKL_PORT" json:"port"`
		RootPassword string `yaml:"root_password" envconfig:"HJKL_ROOT_PASSWORD" json:"root_password"`
	} `yaml:"app" json:"app"`
	Mariadb struct {
		Username string `yaml:"username" envconfig:"HJKL_MARIADB_USERNAME" json:"username"`
		Password string `yaml:"password" envconfig:"HJKL_MARIADB_PASSWORD" json:"password"`
		Addr     string `yaml:"addr" envconfig:"HJKL_MARIADB_ADDR" json:"addr"`
		Db       string `yaml:"db" envconfig:"HJKL_MARIADB_DB" json:"db"`
	} `yaml:"mariadb" json:"mariadb"`
	Redis struct {
		Addr     string `yaml:"addr" envconfig:"HJKL_REDIS_ADDR" json:"addr"`
		Db       int    `yaml:"db" envconfig:"HJKL_REDIS_DB" json:"db"`
		Password string `yaml:"password" envconfig:"HJKL_REDIS_PASSWORD" json:"password"`
	} `yaml:"redis" json:"redis"`
}

func Read() Config {
	once.Do(func() {
		if err := readYAML("config.yaml", &cfg); err != nil {
			panic(err)
		}
		if err := readEnv(&cfg); err != nil {
			panic(err)
		}

		s, _ := json.MarshalIndent(cfg, "# ", "    ")
		fmt.Println("# Config:")
		fmt.Println("# " + string(s))
	})
	return cfg
}

func readYAML(filename string, c *Config) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(c)
}

func readEnv(c *Config) error {
	return envconfig.Process("", c)
}
