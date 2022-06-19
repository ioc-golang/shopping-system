package main

import (
	"context"
	"fmt"
	"github.com/alibaba/ioc-golang"
	conf "github.com/alibaba/ioc-golang/config"
	"github.com/alibaba/ioc-golang/extension/normal/http_server"
	"github.com/alibaba/ioc-golang/extension/normal/http_server/ghttp"
	"github.com/ioc-golang/shopping-system/internal/auth"
	"github.com/ioc-golang/shopping-system/pkg/model/vo"
	"github.com/ioc-golang/shopping-system/pkg/service/festival/api"
	"net/http"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	HttpServer    http_server.ImplIOCInterface   `normal:",port=8080"`
	Authenticator auth.AuthenticatorIOCInterface `singleton:""`

	FestivalServiceClient api.ServiceIOCRPCClient `rpc-client:",address=festival-service"`
}

func (a *App) Run() {
	a.HttpServer.RegisterRouter("/shopping/festival", func(ctx *ghttp.GRegisterController) error {
		req := ctx.Req.(*vo.GetFestivalHomepageRequest)
		if !a.Authenticator.Check(req.UserID) {
			ctx.W.WriteHeader(http.StatusUnauthorized)
			return fmt.Errorf("invalid userID %d", req.UserID)
		}

		cards, totalPage, err := a.FestivalServiceClient.ListCards(req.PageIndex, req.PageSize)
		if err != nil {
			return err
		}

		ctx.Rsp = &vo.GetFestivalHomepageResponse{
			Cards:     cards,
			PageIndex: req.PageIndex,
			TotalPage: totalPage,
		}
		return nil
	}, &vo.GetFestivalHomepageRequest{}, &vo.GetFestivalHomepageResponse{}, http.MethodGet)
	a.HttpServer.Run(context.Background())
}

func main() {
	if err := loadIoC(); err != nil {
		panic(err)
	}

	app, err := GetApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}

func loadIoC() error {
	nameOpt := conf.WithConfigName("ioc_golang")
	typeOpt := conf.WithConfigType("yaml")
	err := ioc.Load(nameOpt, typeOpt)

	return err
}
