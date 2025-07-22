package handler

import (
	"net/http"
	"one-api/router"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// 初始化 One-API 的所有路由规则
	ginRouter := router.SetRouter()
	// 将 Vercel 传来的请求交给 One-API 的路由系统处理
	ginRouter.ServeHTTP(w, r)
}
