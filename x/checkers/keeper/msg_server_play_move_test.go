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

func setupMsgServerWithOneGameForPlayMove(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
  k, ctx := keepertest.CheckersKeeper(t)
  checkers.InitGenesis(ctx, k, *types.DefaultGenesis())
  server := keeper.NewMsgServerImpl(k)
  context := ctx
  server.CreateGame(context, &types.MsgCreateGame{
    Creator: alice,
    Black:   bob,
    Red:     carol,
  })
  return server, k, context
}

func TestPlayMove(t *testing.T) {
  msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
  playMoveResponse, err := msgServer.PlayMove(context, &types.MsgPlayMove{
    Creator:   bob,
    GameIndex: "1",
    FromX:     1,
    FromY:     2,
    ToX:       2,
    ToY:       3,
  })

  require.Nil(t, err)
  require.EqualValues(t, types.MsgPlayMoveResponse{
    CapturedX: -1,
    CapturedY: -1,
    Winner:    "*",
  }, *playMoveResponse)
}

func TestPlayMoveCannotParseGame(t *testing.T) {
  msgServer, k, context := setupMsgServerWithOneGameForPlayMove(t)
  ctx := sdk.UnwrapSDKContext(context)
  storedGame, _ := k.GetStoredGame(ctx, "1")
  storedGame.Board = "not a board"
  k.SetStoredGame(ctx, storedGame)

  defer func() {
    r := recover()
    require.NotNil(t, r, "The code did not panic")
    require.Equal(t, r, "game cannot be parsed: invalid board string: not a board")
  }()

  msgServer.PlayMove(context, &types.MsgPlayMove{
    Creator:   bob,
    GameIndex: "1",
    FromX:     1,
    FromY:     2,
    ToX:       2,
    ToY:       3,
  })
}

func TestPlayMoveEmitted(t *testing.T) {
  msgServer, _, context := setupMsgServerWithOneGameForPlayMove(t)
  msgServer.PlayMove(context, &types.MsgPlayMove{
    Creator:   bob,
    GameIndex: "1",
    FromX:     1,
    FromY:     2,
    ToX:       2,
    ToY:       3,
  })
  ctx := sdk.UnwrapSDKContext(context)
  require.NotNil(t, ctx)
  events := sdk.StringifyEvents(ctx.EventManager().ABCIEvents())
  require.Len(t, events, 2)
  event := events[1]
  require.EqualValues(t, sdk.StringEvent{
    Type: "move-played",
    Attributes: []sdk.Attribute{
      {Key: "creator", Value: bob},
      {Key: "game-index", Value: "1"},
      {Key: "captured-x", Value: "-1"},
      {Key: "captured-y", Value: "-1"},
      {Key: "winner", Value: "*"},
      {Key: "board", Value: "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*"},
    },
  }, event)
}
