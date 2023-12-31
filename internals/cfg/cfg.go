package cfg

import (
	"fmt"
	"github.com/spf13/viper"
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
	v.SetEnvPrefix("SERV") 
	v.SetDefault("PORT", "8080")
	v.SetDefault("DBUSER", "kirill")
	v.SetDefault("DBPASS", "password")
	v.SetDefault("DBHOST", "db")
	v.SetDefault("DBPORT", "5432")
	v.SetDefault("DBNAME", "test")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string { 
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}