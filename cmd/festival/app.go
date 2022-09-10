package main

import (
	"flag"

	"github.com/alibaba/ioc-golang"
	conf "github.com/alibaba/ioc-golang/config"
	"github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"

	_ "github.com/ioc-golang/shopping-system/pkg/service/festival"
)

func main() {
	var mode = flag.String("m", "local", "which profile to be activated, support k8s, docker, local")
	flag.Parse()

	if *mode == "local" {
		// run shopping-system locally, change ioc rpc server port to solve conflict (default is :2022).
		rpc_service.SetParam(&rpc_service.Param{ExportPort: "2024"})
	}

	if err := ioc.Load(
		conf.WithConfigName("festival"),
		conf.WithProfilesActive(*mode)); err != nil {
		panic(err)
	}
	select {}
}
