package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/vonum/checkers/x/checkers/rules"
	"github.com/vonum/checkers/x/checkers/testutil"
	"github.com/vonum/checkers/x/checkers/types"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() types.StoredGame {
	return types.StoredGame{
		Black: alice,
		Red:   bob,
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
	}
}

func TestGetBlackAddress(t *testing.T) {
	address, _ := sdk.AccAddressFromBech32(alice)
	blackAddress, err := GetStoredGame1().GetBlackAddress()

	require.EqualValues(t, address, blackAddress)
	require.Nil(t, err)
}

func TestGetBlackAddressInvalid(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"

	blackAddress, err := storedGame.GetBlackAddress()
	require.Nil(t, blackAddress)
	require.EqualError(
		t,
		err,
		"black address is invalid cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)",
	)
	require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestGetRedAddress(t *testing.T) {
	address, _ := sdk.AccAddressFromBech32(bob)
	redAddress, err := GetStoredGame1().GetRedAddress()

	require.EqualValues(t, address, redAddress)
	require.Nil(t, err)
}

func TestGetRedAddressInvalid(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Red = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4"

	redAddress, err := storedGame.GetRedAddress()
	require.Nil(t, redAddress)
	require.EqualError(
		t,
		err,
		"red address is invalid cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)",
	)
	require.EqualError(t, storedGame.Validate(), err.Error())
}
func TestGameValidateOk(t *testing.T) {
	storedGame := GetStoredGame1()
	require.NoError(t, storedGame.Validate())
}
