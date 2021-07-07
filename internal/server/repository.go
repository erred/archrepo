package server

import (
	"context"

	archrepo "go.seankhliao.com/apis/seankhliao/archrepo/v1alpha2"
	"go.seankhliao.com/archrepo/internal/filesystem"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateRepository(ctx context.Context, req *archrepo.CreateRepositoryRequest) (*archrepo.Repository, error) {
	err := s.store.UpdateRepository(ctx, filesystem.Repository{
		Name: req.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &archrepo.Repository{
		Name: req.Name,
	}, nil
}

func (s *Server) DeleteRepository(ctx context.Context, req *archrepo.DeleteRepositoryRequest) (*emptypb.Empty, error) {
	err := s.store.DeleteRepository(ctx, filesystem.Repository{
		Name: req.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

// func (s *Server) GetRepository(context.Context, *archrepo.GetRepositoryRequest) (*archrepo.Repository, error)

func (s *Server) ListRepositories(ctx context.Context, req *archrepo.ListRepositoriesRequest) (*archrepo.ListRepositoriesResponse, error) {
	repos, err := s.store.ListRepositories(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var res archrepo.ListRepositoriesResponse
	for _, repo := range repos {
		res.Repositories = append(res.Repositories, &archrepo.Repository{
			Name: repo.Name,
		})
	}

	return &res, nil
}

func (s *Server) UpdateRepository(ctx context.Context, req *archrepo.UpdateRepositoryRequest) (*archrepo.Repository, error) {
	err := s.store.UpdateRepository(ctx, filesystem.Repository{
		Name: req.Repository.Name,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &archrepo.Repository{
		Name: req.Repository.Name,
	}, nil
}
