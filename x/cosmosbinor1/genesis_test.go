package cosmosbinor1_test

import (
	"testing"

	keepertest "github.com/binordev/cosmos-binor1/testutil/keeper"
	"github.com/binordev/cosmos-binor1/x/cosmosbinor1"
	"github.com/binordev/cosmos-binor1/x/cosmosbinor1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Cosmosbinor1Keeper(t)
	cosmosbinor1.InitGenesis(ctx, *k, genesisState)
	got := cosmosbinor1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
