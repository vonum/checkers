package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vonum/checkers/testutil/keeper"
	"github.com/vonum/checkers/testutil/nullify"
	"github.com/vonum/checkers/x/checkers/keeper"
	"github.com/vonum/checkers/x/checkers/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredGame(keeper keeper.Keeper, ctx context.Context, n int) []types.StoredGame {
	items := make([]types.StoredGame, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredGame(ctx, items[i])
	}
	return items
}

func TestStoredGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredGame(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredGame(ctx)),
	)
}
