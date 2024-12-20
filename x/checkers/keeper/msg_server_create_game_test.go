package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/vonum/checkers/testutil/keeper"
	"github.com/vonum/checkers/x/checkers/keeper"
	checkers "github.com/vonum/checkers/x/checkers/module"
	"github.com/vonum/checkers/x/checkers/types"
)

func setupMsgServerCreateGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {

	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(k), k, ctx
}

func TestCreateGame(t *testing.T) {
	msgServer, _, context := setupMsgServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(
		context,
		&types.MsgCreateGame{
			Creator: alice,
			Black:   bob,
			Red:     alice,
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
  ctx := sdk.UnwrapSDKContext(context)
	systemInfo, found := keeper.GetSystemInfo(ctx)
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
    Winner: "*",
    Deadline: types.FormatDeadline(ctx.BlockTime().Add(types.MaxTurnDuration)),
    MoveCount: 0,
	}, game1)
}

func TestCreate1GameEmitted(t *testing.T) {
  msgSrvr, _, context := setupMsgServerCreateGame(t)
  msgSrvr.CreateGame(context, &types.MsgCreateGame{
    Creator: alice,
    Black:   bob,
    Red:     carol,
  })
  ctx := sdk.UnwrapSDKContext(context)
  require.NotNil(t, ctx)
  events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
  require.Len(t, events, 1)
  event := events[0]
  require.EqualValues(t, sdk.StringEvent{
    Type: "new-game-created",
    Attributes: []sdk.Attribute{
      {Key: "creator", Value: alice},
      {Key: "game-index", Value: "1"},
      {Key: "black", Value: bob},
      {Key: "red", Value: carol},
    },
  }, event)
}
