package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	tracerProvider, err := newTracerProvider()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)
	return
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTracerProvider() (*trace.TracerProvider, error) {
	traceExporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint("172.17.0.1:4318"),
	)
	if err != nil {
		return nil, err
	}
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			traceExporter,
			trace.WithBatchTimeout(time.Second),
		),
	)
	return tracerProvider, nil
}

func main() {
	log.Println("started")
	defer log.Println("exited")
	endpoint := flag.String("endpoint", "", "endpoint to call")
	flag.Parse()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		panic(err)
	}
	defer otelShutdown(context.Background())
	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Printf("received request with headers %v", r.Header)
		if *endpoint != "" {
			log.Println("sending request to " + *endpoint)
			r, _ := http.NewRequestWithContext(ctx, "GET", *endpoint, nil)
			client := http.Client{
				Transport: otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(otel.GetTracerProvider())),
			}
			rs, _ := client.Do(r)
			b, _ := io.ReadAll(rs.Body)
			fmt.Println(string(b))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})
	handler = otelhttp.NewHandler(handler, "", otelhttp.WithTracerProvider(otel.GetTracerProvider()))
	if err := http.ListenAndServe(":80", handler); err != nil {
		panic(err)
	}
}
