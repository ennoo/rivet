package main

import (
	"fmt"
	"github.com/ennoo/rivet/rivet"
	"github.com/ennoo/rivet/utils/tracer/examples"
	"github.com/gin-gonic/gin"
	"github.com/openzipkin/zipkin-go"
)

func main() {

	rivet.Initialize(false, false, false);

	//tracer := examples.GetTracer("demoService", "127.0.01")


	rivet.ListenAndServe(&rivet.ListenServe{
		Engine:      rivet.SetupRouter(testRouter1),
		DefaultPort: "8081",
	})

	//
	//serverMiddleware := zipkinhttp.NewServerMiddleware(
	//	tracer, zipkinhttp.TagResponseSize(true),
	//)
	//
	//// create global zipkin traced http client
	//client, err := zipkinhttp.NewClient(tracer, zipkinhttp.ClientTrace(true))
	//if err != nil {
	//	log.Fatalf("unable to create client: %+v\n", err)
	//}


	//ts := httptest.NewServer(serverMiddleware());

}



func testRouter1(engine *gin.Engine) {
	// 仓库相关路由设置
	vRepo := engine.Group("/rivet")
	vRepo.GET("/get",func(context *gin.Context) {
		tracer  := examples.GetTracer("feifeiDemoService","172.20.23.100:80");
		// tracer can now be used to create spans.
		span := tracer.StartSpan("first_soperation")




		fmt.Println("hehesda");

		// ... do some work ...
		span.SetName("feifeis");
		span.Finish()

		childSpan := tracer.StartSpan("some_operation2", zipkin.Parent(span.Context()))

		//time.Sleep(10000);
		//... do some work ...
		childSpan.Finish()

		span.Finish()
	})

}

