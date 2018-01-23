	# On main.go add

    ```go
    //zipkin with transport
	transport, err := zipkin.NewHTTPTransport(
		"http://localhost:9411/api/v1/spans",
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	if err != nil {
		log.Fatalf("Cannot initialize HTTP transport: %v", err)
	}

	tracer, closer := jaeger.NewTracer(
		"server",
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(transport, nil),
	)
	defer closer.Close()

    opentracing.InitGlobalTracer(tracer)
    ```