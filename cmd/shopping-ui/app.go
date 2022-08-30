package main

import (
	"flag"
	"net/http"

	"github.com/alibaba/ioc-golang"
	"github.com/alibaba/ioc-golang/config"
	"github.com/gin-gonic/gin"

	"github.com/ioc-golang/shopping-system/internal/auth"
	"github.com/ioc-golang/shopping-system/pkg/model/vo"
	festivalAPI "github.com/ioc-golang/shopping-system/pkg/service/festival/api"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	Authenticator auth.AuthenticatorIOCInterface `singleton:""`

	FestivalServiceClient festivalAPI.ServiceIOCRPCClient `rpc-client:""`
}

func (a *App) Run() {
	engion := gin.Default()
	// biz http handler function
	engion.GET("/festival/listCards", func(c *gin.Context) {
		req := &vo.GetFestivalHomepageRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !a.Authenticator.Check(req.UserID) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		cards, err := a.FestivalServiceClient.ListCards(req.UserID, req.Num)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.PureJSON(http.StatusOK, &vo.GetFestivalHomepageResponse{
			Cards: cards,
		})
		return
	})
	engion.GET("/festival/listCachedCards", func(c *gin.Context) {
		req := &vo.GetFestivalHomepageRequest{}
		if err := c.ShouldBindQuery(req); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !a.Authenticator.Check(req.UserID) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		cards, err := a.FestivalServiceClient.ListCachedCards(req.UserID, req.Num)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.PureJSON(http.StatusOK, &vo.GetFestivalHomepageResponse{
			Cards: cards,
		})
		return
	})

	if err := engion.Run(":8080"); err != nil {
		panic(err)
	}
}

func main() {
	var mode = flag.String("m", "local", "which profile to be activated, support k8s, docker, local")
	flag.Parse()

	if err := ioc.Load(
		config.WithConfigName("shopping_ui"),
		config.WithProfilesActive(*mode)); err != nil {
		panic(err)
	}

	app, err := GetAppSingleton()
	if err != nil {
		panic(err)
	}
	app.Run()
}
