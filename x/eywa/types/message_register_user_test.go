package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/eywa-foundation/eywa-go-client/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterUser
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterUser{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgRegisterUser{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
