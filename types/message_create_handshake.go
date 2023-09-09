package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateHandshake = "create_handshake"

var _ sdk.Msg = &MsgCreateHandshake{}

func NewMsgCreateHandshake(creator string, receiver string, roomId string, serverAddress string) *MsgCreateHandshake {
	return &MsgCreateHandshake{
		Creator:       creator,
		Receiver:      receiver,
		RoomId:        roomId,
		ServerAddress: serverAddress,
	}
}

func (msg *MsgCreateHandshake) Route() string {
	return RouterKey
}

func (msg *MsgCreateHandshake) Type() string {
	return TypeMsgCreateHandshake
}

func (msg *MsgCreateHandshake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateHandshake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateHandshake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
