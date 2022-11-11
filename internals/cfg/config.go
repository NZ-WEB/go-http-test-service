package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Cfg struct {
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()
	v.SetEnvPrefix("SERF")
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "postgres")
	v.SetDefault("DBPASS", "example")
	v.SetDefault("DBHOST", "localhost")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "postgres")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}

	return cfg
}

func (cfg *Cfg) GetDbString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
