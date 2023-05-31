package voter_test

import (
	"testing"

	keepertest "github.com/fanfury-sports/blackfury/testutil/keeper"
	"github.com/fanfury-sports/blackfury/testutil/nullify"
	"github.com/fanfury-sports/blackfury/x/voter"
	"github.com/fanfury-sports/blackfury/x/voter/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VoterKeeper(t)
	voter.InitGenesis(ctx, *k, genesisState)
	got := voter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
