package keeper

import (
	"github.com/nzee/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
