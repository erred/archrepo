package server

import (
	"context"

	archrepo "go.seankhliao.com/apis/seankhliao/archrepo/v1alpha2"
	"go.seankhliao.com/archrepo/internal/filesystem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreatePackageVersion(ctx context.Context, req *archrepo.CreatePackageVersionRequest) (*archrepo.PackageVersion, error) {
	err := s.store.UpdatePackageVersion(ctx, filesystem.Repository{
		Name: req.Parent,
	}, filesystem.PackageVersion{
		Name: req.Name,
		Data: req.Data,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &archrepo.PackageVersion{Name: req.Name}, nil
}

func (s *Server) DeletePackageVersion(ctx context.Context, req *archrepo.DeletePackageVersionRequest) (*emptypb.Empty, error) {
	err := s.store.DeletePackageVersion(ctx, filesystem.Repository{
		Name: req.Parent,
	}, filesystem.PackageVersion{
		Name: req.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// func (s *Server) GetPackageVersion(context.Context, *archrepo.GetPackageVersionRequest) (*archrepo.PackageVersion, error)
// func (s *Server) ListPackageVersions(context.Context, *archrepo.ListPackageVersionsRequest) (*archrepo.ListPackageVersionsResponse, error)

func (s *Server) UpdatePackageVersion(ctx context.Context, req *archrepo.UpdatePackageVersionRequest) (*archrepo.PackageVersion, error) {
	err := s.store.UpdatePackageVersion(ctx, filesystem.Repository{
		Name: req.Repository.Parent,
	}, filesystem.PackageVersion{
		Name: req.Repository.Name,
		Data: req.Repository.Data,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &archrepo.PackageVersion{Name: req.Repository.Name}, nil
}
