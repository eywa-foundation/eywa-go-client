package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterUser = "register_user"

var _ sdk.Msg = &MsgRegisterUser{}

func NewMsgRegisterUser(creator string, pubkey string) *MsgRegisterUser {
	return &MsgRegisterUser{
		Creator: creator,
		Pubkey:  pubkey,
	}
}

func (msg *MsgRegisterUser) Route() string {
	return RouterKey
}

func (msg *MsgRegisterUser) Type() string {
	return TypeMsgRegisterUser
}

func (msg *MsgRegisterUser) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterUser) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
