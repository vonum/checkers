package checkers_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vonum/checkers/testutil/keeper"
	"github.com/vonum/checkers/testutil/nullify"
	checkers "github.com/vonum/checkers/x/checkers/module"
	"github.com/vonum/checkers/x/checkers/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: &types.SystemInfo{
			NextId: 60,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, k, genesisState)
	got := checkers.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	// this line is used by starport scaffolding # genesis/test/assert
}
