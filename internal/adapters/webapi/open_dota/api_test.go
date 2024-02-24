package opendota_test

import (
	"context"
	"testing"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	"github.com/stretchr/testify/require"
)

var key = "f0ebafea-348b-48e1-a0c9-134a9b88211b"

func TestAPI(t *testing.T) {
	r := require.New(t)
	api := opendota.New(key)

	patches, err := api.Patches(context.TODO())
	r.NoError(err)
	r.NotEmpty(patches)
}
