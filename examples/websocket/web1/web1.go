package main

import (
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/log"
)

func main() {
	rivet.Initialize(false, false, false)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter1),
		DefaultPort: "8081",
	})
}

func testRouter1(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/ws/:id", rivet.WS)
	router.GET("/send", send)
}

func send(router *response.Router) {
	//context := router.Context
	//context.Header("Upgrade", "websocket")
	//rivet.Request().Call(router.Context, http.MethodGet, "http://localhost:8081", "rivet/ws")
	for index := range rivet.Keepers {
		if err := rivet.Keepers[index].Write([]byte("hello")); nil != err {
			log.Self.Warn("err", log.Error(err))
		}
	}
}
