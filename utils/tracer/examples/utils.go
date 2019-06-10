package examples

import (
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
)

const (
	// 运行前需要修改endpointUrl的值，从https://tracing-analysis.console.aliyun.com/ 获取zipkin网关
	enpoitUrl = "http://test.hbaas.vh.cn:9411/api/v2/spans"
)

func GetTracer(serviceName string, ip string) *zipkin.Tracer {

	// create a reporter to be used by the tracer
	reporter := httpreporter.NewReporter(enpoitUrl)

	// set-up the local endpoint for our service
	endpoint, _ := zipkin.NewEndpoint(serviceName, ip)

	// set-up our sampling strategy
	sampler := zipkin.NewModuloSampler(1)

	// initialize the tracer
	tracer, _ := zipkin.NewTracer(
		reporter,
		zipkin.WithLocalEndpoint(endpoint),
		zipkin.WithSampler(sampler),
	)
	return tracer;
}
