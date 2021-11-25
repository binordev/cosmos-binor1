package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/binordev/cosmos-binor1/testutil/keeper"
	"github.com/binordev/cosmos-binor1/x/cosmosbinor1/keeper"
	"github.com/binordev/cosmos-binor1/x/cosmosbinor1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Cosmosbinor1Keeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
