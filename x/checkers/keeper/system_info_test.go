package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/vonum/checkers/testutil/keeper"
	"github.com/vonum/checkers/testutil/nullify"
	"github.com/vonum/checkers/x/checkers/keeper"
	"github.com/vonum/checkers/x/checkers/types"
)

func createTestSystemInfo(keeper keeper.Keeper, ctx context.Context) types.SystemInfo {
	item := types.SystemInfo{}
	keeper.SetSystemInfo(ctx, item)
	return item
}

func TestSystemInfoGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	item := createTestSystemInfo(keeper, ctx)
	rst, found := keeper.GetSystemInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestSystemInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	createTestSystemInfo(keeper, ctx)
	keeper.RemoveSystemInfo(ctx)
	_, found := keeper.GetSystemInfo(ctx)
	require.False(t, found)
}
