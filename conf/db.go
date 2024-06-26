package conf

import (
	"fmt"
	"sync"

	// "github.com/hpp131/gblog/ioc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const APPNAME = "database"
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

// func init() {
// 	ioc.Config().Registry(APPNAME)
// }


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

// default config
func defaultConfig() *Config{
	return &Config{
		Mysql: &Mysql{Username: "blog",
			Password: "blog111",
			Host:     "192.168.0.34",
			Port:     30007,
			DB:       "blogs",
			Debug:    true,},
	}
}

func (c *Config) DB() *gorm.DB{
	return c.Mysql.Get()
}

