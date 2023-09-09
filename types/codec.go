package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterUser{}, "eywa/RegisterUser", nil)
	cdc.RegisterConcrete(&MsgCreateHandshake{}, "eywa/CreateHandshake", nil)
	cdc.RegisterConcrete(&MsgCreateRelayServer{}, "eywa/CreateRelayServer", nil)
	cdc.RegisterConcrete(&MsgCreateChat{}, "eywa/CreateChat", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterUser{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateHandshake{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRelayServer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChat{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
