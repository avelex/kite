package opendota_test

import (
	"context"
	"fmt"
	"testing"

	opendota "github.com/avelex/kite/internal/adapters/webapi/open_dota"
	"github.com/stretchr/testify/require"
)

var key = "f0ebafea-348b-48e1-a0c9-134a9b88211b"

func TestAPI(t *testing.T) {
	r := require.New(t)
	api := opendota.New(key)

	m, err := api.PlayerAllMatches(context.Background(), 18180970)
	r.NoError(err)

	match := m[0]

	p, err := api.Match(context.Background(), match.MatchID)
	r.NoError(err)

	f := p.Player(18180970)
	fmt.Printf("p: %+v\n", f.Obs)
}
