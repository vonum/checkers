package types

import (
	"fmt"

	"github.com/vonum/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
)

func (storedGame StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
  black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
  return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame StoredGame) GetRedAddress() (black sdk.AccAddress, err error) {
  black, errRed := sdk.AccAddressFromBech32(storedGame.Red)
  return black, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame StoredGame) ParseGame() (game *rules.Game, err error) {
  board, errBoard := rules.Parse(storedGame.Board)
  if errBoard != nil {
      return nil, sdkerrors.Wrapf(errBoard, ErrGameNotParseable.Error())
  }
  board.Turn = rules.StringPieces[storedGame.Turn].Player
  if board.Turn.Color == "" {
      return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
  }
  return board, nil
}

func (storedGame StoredGame) Validate() (err error) {
  _, err = storedGame.GetBlackAddress()
  if err != nil {
      return err
  }
  _, err = storedGame.GetRedAddress()
  if err != nil {
      return err
  }
  _, err = storedGame.ParseGame()
  return err
}
