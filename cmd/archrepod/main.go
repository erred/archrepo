package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"go.seankhliao.com/archrepo/internal/server"
	"go.seankhliao.com/w/v16/webserver"
	"k8s.io/klog/v2/klogr"
)

func main() {
	so := server.NewOptions(flag.CommandLine)
	wo := webserver.NewOptions(flag.CommandLine)
	flag.Parse()

	ctx := context.Background()
	ctx, _ = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)

	l := klogr.New()

	m, err := server.New(ctx, so)
	if err != nil {
		l.Error(err, "setup")
		os.Exit(1)
	}

	wo.Logger = l
	wo.Handler = m

	webserver.New(ctx, wo).Run(ctx)
}
