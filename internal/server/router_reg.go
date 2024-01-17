package server

import (
	"github.com/Godyu97/vege9/midware"
	"github.com/Godyu97/vege9/vegeRouter"
	_ "github.com/Godyu97/vegeGoWebTemplate/docs"
	"github.com/Godyu97/vegeGoWebTemplate/internal/handler"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegHttpRouter(mux *gin.Engine) {
	mux.Use(
		gin.Logger(),
		gin.Recovery(),
		midware.Cors(),
	)
	//ping
	mux.GET("/Ping", handler.Ping)
	//pprof
	pprof.Register(mux)
	//apis
	api := mux.Group("/v1/api")
	//obj
	obj := api.Group("/obj")
	obj.GET(vegeRouter.PathUri, vegeRouter.RegApiHandler(GetApiObj))
	obj.POST(vegeRouter.PathUri, vegeRouter.RegApiHandler(PostApiObj))
	//file server
	api.Static("/static", "./static")
	//swag docs
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

var UriToFnNameM = map[string]string{}
var GetApiObj = handler.GetApi{UriToFnNameM: UriToFnNameM}
var PostApiObj = handler.PostApi{UriToFnNameM: UriToFnNameM}
