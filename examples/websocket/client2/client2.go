package main

import (
	"github.com/ennoo/rivet"
	"github.com/ennoo/rivet/trans/response"
	"github.com/ennoo/rivet/utils/log"
	"golang.org/x/net/websocket"
)

func main() {
	rivet.Initialize(false, false, false)
	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter1),
		DefaultPort: "8083",
	})
}

func testRouter1(router *response.Router) {
	// 仓库相关路由设置
	router.Group = router.Engine.Group("/rivet")
	router.GET("/get", get)
	router.GET("/send", send)
	router.GET("/close", close)
}

func get(router *response.Router) {
	//context := router.Context
	//context.Header("Upgrade", "websocket")
	//rivet.Request().Call(router.Context, http.MethodGet, "http://localhost:8081", "rivet/ws")
	connect()
	sendAndReceive()
}

func send(router *response.Router) {
	//context := router.Context
	//context.Header("Upgrade", "websocket")
	//rivet.Request().Call(router.Context, http.MethodGet, "http://localhost:8081", "rivet/ws")
	sendAndReceive()
}

func close(router *response.Router) {
	closeConn()
}

var (
	origin = "http://localhost:8081/"
	url    = "ws://localhost:8081/rivet/ws/young2"
	conn   *websocket.Conn
	err    error
)

func connect() {
	conn, err = websocket.Dial(url, "", origin)
	if err != nil {
		log.Self.Warn("error", log.Error(err))
	}
	go receive()
}

func sendAndReceive() {
	message := []byte("hello, world!你好")
	_, err = conn.Write(message)
	if err != nil {
		log.Self.Warn("error", log.Error(err))
	}
	log.Self.Debug("Send", log.String("message", string(message)))
}

func closeConn() {
	conn.Close() //关闭连接
}

func receive() {
	for {
		var msg = make([]byte, 512)
		m, err := conn.Read(msg)
		if err != nil {
			log.Self.Warn("error", log.Error(err))
		}
		log.Self.Debug("Receive", log.String("message", string(msg[:m])))
	}
}
