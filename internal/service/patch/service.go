package patch

import (
	"context"
	"time"
)

type Repository interface {
	LatestPatchVersion(ctx context.Context) (int64, error)
	PatchVersionFromString(ctx context.Context, patch string) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) LatestPatchVersion(ctx context.Context) int64 {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	version, err := s.repo.LatestPatchVersion(ctxTimeout)
	if err != nil {
		return 0
	}

	return version
}

func (s *service) PatchVersionFromString(ctx context.Context, patch string) int64 {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	version, err := s.repo.PatchVersionFromString(ctxTimeout, patch)
	if err != nil {
		return 0
	}

	return version
}
