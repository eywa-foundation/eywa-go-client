package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/eywa-foundation/eywa-go-client/x/eywa/keeper"
	"github.com/eywa-foundation/eywa-go-client/x/eywa/types"
)

func SimulateMsgRegisterUser(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRegisterUser{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RegisterUser simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RegisterUser simulation not implemented"), nil, nil
	}
}
