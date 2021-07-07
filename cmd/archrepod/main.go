package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"go.seankhliao.com/archrepo/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var o server.Options
	var haddr, gaddr string
	flag.StringVar(&haddr, "http-addr", ":8080", "http address")
	flag.StringVar(&gaddr, "grpc-addr", ":8081", "grpc address")
	o.InitFlags(flag.CommandLine)
	flag.Parse()

	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	arsvr, err := server.New(ctx, o)
	if err != nil {
		log.Println(err)
		return
	}

	svr := grpc.NewServer()
	arsvr.Register(svr)
	reflection.Register(svr)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		log.Println("grpc listening on", gaddr)
		lis, err := net.Listen("tcp", gaddr)
		if err != nil {
			log.Println(err)
			cancel()
			return
		}
		err = svr.Serve(lis)
		if err != nil {
			log.Println(err)
			cancel()
			return
		}
	}()

	hsvr := &http.Server{
		Addr:    haddr,
		Handler: arsvr,
	}
	go func() {
		defer wg.Done()

		log.Println("http listening on", hsvr.Addr)
		err := hsvr.ListenAndServe()
		if err != nil {
			log.Println(err)
			cancel()
			return
		}
	}()

	<-ctx.Done()
	go hsvr.Shutdown(context.Background())
	go svr.GracefulStop()

	wg.Wait()
}
