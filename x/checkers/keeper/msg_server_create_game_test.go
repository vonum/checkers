package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	checkers "github.com/vonum/checkers/x/checkers/module"
	keepertest "github.com/vonum/checkers/testutil/keeper"
	"github.com/stretchr/testify/require"
	"github.com/vonum/checkers/x/checkers/keeper"
	"github.com/vonum/checkers/x/checkers/types"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {

    k, ctx := keepertest.CheckersKeeper(t)
    checkers.InitGenesis(ctx, k, *types.DefaultGenesis())
    return keeper.NewMsgServerImpl(k), k, ctx
}

func TestCreateGame(t *testing.T) {
  msgServer, _, context :=  setupMsgServerCreateGame(t)
  createResponse, err := msgServer.CreateGame(
    context,
    &types.MsgCreateGame{
      Creator: alice,
      Black: bob,
      Red: alice,
    },
  )

  require.Nil(t, err)
  require.EqualValues(
    t,
    types.MsgCreateGameResponse{
      GameIndex: "1",
    },
    *createResponse,
  )
}

func TestCreate1GameHasSaved(t *testing.T) {
  msgSrvr, keeper, context := setupMsgServerCreateGame(t)
  msgSrvr.CreateGame(context, &types.MsgCreateGame{
    Creator: alice,
    Black:   bob,
    Red:     carol,
  })
  systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
  require.True(t, found)
  require.EqualValues(t, types.SystemInfo{
      NextId: 2,
  }, systemInfo)
  game1, found1 := keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
  require.True(t, found1)
  require.EqualValues(t, types.StoredGame{
    Index: "1",
    Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
    Turn:  "b",
    Black: bob,
    Red:   carol,
  }, game1)
}
