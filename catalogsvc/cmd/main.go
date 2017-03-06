package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/mvonbodun/geaux-commerce/catalogsvc"
	"github.com/mvonbodun/geaux-commerce/catalogsvc/category"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var s catalogsvc.Service
	{
		s = catalogsvc.NewService(category.Repo)
	}

	var h http.Handler
	{
		h = catalogsvc.MakeHTTPHandler(ctx, s, log.NewContext(logger).With("component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}
