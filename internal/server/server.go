package server

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	archrepo "go.seankhliao.com/apis/seankhliao/archrepo/v1alpha2"
	"go.seankhliao.com/archrepo/internal/filesystem"
	"google.golang.org/grpc"
)

type Options struct {
	Root string
}

func (o *Options) InitFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Root, "data", "/opt/archrepo", "path to root directory")
}

type Server struct {
	store      *filesystem.Store
	fileServer http.Handler

	archrepo.UnimplementedRepositoryServiceServer
}

func New(ctx context.Context, opt Options) (*Server, error) {
	store, err := filesystem.New(ctx, opt.Root)
	if err != nil {
		return nil, fmt.Errorf("create backing store: %w", err)
	}

	fileServer := http.FileServer(http.FS(os.DirFS(filepath.Join(opt.Root, "repos"))))

	return &Server{
		store:      store,
		fileServer: fileServer,
	}, nil
}

func (s *Server) Register(svr *grpc.Server) {
	archrepo.RegisterRepositoryServiceServer(svr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.fileServer.ServeHTTP(w, r)
}
