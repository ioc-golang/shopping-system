package main

import (
	"flag"
	"io/ioutil"
	"strings"

	"github.com/alibaba/ioc-golang"
	conf "github.com/alibaba/ioc-golang/config"
	"github.com/alibaba/ioc-golang/extension/config"
	"github.com/alibaba/ioc-golang/extension/db/gorm"

	_ "github.com/ioc-golang/shopping-system/pkg/service/product"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type MysqlDBInitializer struct {
	GORMDB          gorm.GORMDBIOCInterface `normal:",dev-mysql"`
	InitSQLFilePath *config.ConfigString    `config:",config.initSqlFilePath"`
}

func (m *MysqlDBInitializer) Init() {
	if data, err := ioutil.ReadFile(m.InitSQLFilePath.Value()); err != nil {
		panic(err)
	} else {
		m.GORMDB.Exec("drop schema dev;")
		m.GORMDB.Exec("create schema dev;")
		for _, sql := range strings.Split(string(data), ";") {
			m.GORMDB.Exec(sql + ";")
		}
	}
}

func main() {
	var mode = flag.String("m", "local", "which profile to be activated, support k8s, docker, local")
	flag.Parse()

	if err := ioc.Load(
		conf.WithConfigName("product"),
		conf.WithProfilesActive(*mode)); err != nil {
		panic(err)
	}

	initializer, err := GetMysqlDBInitializerIOCInterfaceSingleton()
	if err != nil {
		panic(err)
	}
	initializer.Init()
	select {}
}
