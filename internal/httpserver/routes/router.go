package routes

import (
	"gin-scaffold/internal/httpserver/appcontext"
	"gin-scaffold/internal/httpserver/handlers"
	"github.com/gin-gonic/gin"
)

// Register 函数注册 http 路由
// 可在此函数中设置模板和设置静态文件路径
func Register(router *gin.Engine, appCtx *appcontext.Context) {
	router.GET("/", handlers.Welcome(appCtx))
}
