package conf

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var config *Config

type Config struct{
	Mysql  *Mysql
}

type Mysql struct{
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`
	// 判断这个私有属性, 来判断是否返回已有的对象
	db *gorm.DB
	l  sync.Mutex

}

func (m *Mysql) DSN() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	m.Username,
	m.Password,
	m.Host,
	m.Port,
	m.DB,
	)
}

func (m *Mysql) Get() *gorm.DB{
	m.l.Lock()
	defer m.l.Unlock()

	if m.db == nil{
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err!= nil{
			panic(err)
		}
		m.db = db
	}
	return m.db
}

func C() *Config{
	if config == nil{
		return defaultConfig()
	}
	return config
}

func defaultConfig() *Config{
	return &Config{
		Mysql: &Mysql{Username: "root",
			Password: "Strong@01",
			Host:     "123.60.159.166",
			Port:     3306,
			DB:       "blogs",
			Debug:    true,},
	}
}

func (c *Config) DB() *gorm.DB{
	return c.Mysql.Get()
}