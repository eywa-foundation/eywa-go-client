package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateRelayServer = "create_relay_server"

var _ sdk.Msg = &MsgCreateRelayServer{}

func NewMsgCreateRelayServer(creator string, nickname string, location string) *MsgCreateRelayServer {
	return &MsgCreateRelayServer{
		Creator:  creator,
		Nickname: nickname,
		Location: location,
	}
}

func (msg *MsgCreateRelayServer) Route() string {
	return RouterKey
}

func (msg *MsgCreateRelayServer) Type() string {
	return TypeMsgCreateRelayServer
}

func (msg *MsgCreateRelayServer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRelayServer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRelayServer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
