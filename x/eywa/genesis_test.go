package eywa_test

import (
	"testing"

	keepertest "github.com/eywa-foundation/eywa-go-client/testutil/keeper"
	"github.com/eywa-foundation/eywa-go-client/testutil/nullify"
	"github.com/eywa-foundation/eywa-go-client/x/eywa"
	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EywaKeeper(t)
	eywa.InitGenesis(ctx, *k, genesisState)
	got := eywa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
