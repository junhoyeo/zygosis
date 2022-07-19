package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "zygosis/testutil/keeper"
	"zygosis/x/zygosis/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ZygosisKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
