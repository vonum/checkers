package keeper

import (
	"github.com/vonum/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
