package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysRoleRouter)
}

// 默认需登录认证的路由
func registerSysRoleRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-role").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiSysRole.Get)
		r.POST("/create", apis.ApiSysRole.Create)
		r.POST("/update", apis.ApiSysRole.Update)
		r.POST("/page", apis.ApiSysRole.QueryPage)
		r.POST("/del", apis.ApiSysRole.Del)
	}
}