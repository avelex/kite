package memory

import (
	"context"
	"time"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
)

type repository struct {
	openDotaAPI  opendota.API
	patchVersion map[string]int64
	versionPatch map[int64]string
	latestPatch  string
}

func New(open opendota.API) *repository {
	return &repository{
		openDotaAPI:  open,
		patchVersion: make(map[string]int64),
		versionPatch: make(map[int64]string),
	}
}

func (r *repository) PrepareData(ctx context.Context) error {
	ctxPatch, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	patches, err := r.openDotaAPI.Patches(ctxPatch)
	if err != nil {
		return err
	}

	for _, p := range patches {
		r.patchVersion[p.Name] = p.ID
		r.versionPatch[p.ID] = p.Name
	}

	if len(patches) > 0 {
		r.latestPatch = patches[len(patches)-1].Name
	}

	return nil
}

func (r *repository) LatestPatchVersion(_ context.Context) (int64, error) {
	version, ok := r.patchVersion[r.latestPatch]
	if !ok {
		return -1, nil
	}

	return version, nil
}

func (r *repository) PatchVersionFromString(_ context.Context, patch string) (int64, error) {
	version, ok := r.patchVersion[patch]
	if !ok {
		return -1, nil
	}

	return version, nil
}
