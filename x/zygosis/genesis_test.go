package zygosis_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "zygosis/testutil/keeper"
	"zygosis/testutil/nullify"
	"zygosis/x/zygosis"
	"zygosis/x/zygosis/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZygosisKeeper(t)
	zygosis.InitGenesis(ctx, *k, genesisState)
	got := zygosis.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
