package router

import "github.com/gin-gonic/gin"

type (
	Engine      = gin.Engine
	RoutesInfo  = gin.RoutesInfo
	RouteInfo   = gin.RouteInfo
	RouterGroup = gin.RouterGroup
)

var New = gin.New
