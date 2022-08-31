package migrations

import (
	"io/ioutil"

	"github.com/Moonieh/banking/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
}

func GetConf() *Config {
	mapconf := MapConf()
	conf := SetConfig(mapconf)

	return conf
}

func MapConf() map[string]string {
	read, err := ioutil.ReadFile("conf/conf.yaml")
	helpers.HandleErr(err)
	c := make(map[string]string)
	err = yaml.Unmarshal(read, c)
	helpers.HandleErr(err)
	return c
}

func SetConfig(m map[string]string) *Config {
	c := &Config{
		Username: m["username"],
		Password: m["password"],
		Host:     m["host"],
		Port:     m["port"],
	}

	return c
}

func ConnectDB(c *Config) *sqlx.DB {
	db, err := sqlx.Connect("mysql", c.Username+":"+c.Password+"@("+c.Host+":"+c.Port+")")
	helpers.HandleErr(err)
	return db
}
