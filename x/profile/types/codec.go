package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProfile{}, "profile/CreateProfile", nil)
	cdc.RegisterConcrete(&MsgUpdateProfile{}, "profile/UpdateProfile", nil)
	cdc.RegisterConcrete(&MsgDeleteProfile{}, "profile/DeleteProfile", nil)
	cdc.RegisterConcrete(&MsgAddBioToProfile{}, "profile/AddBioToProfile", nil)
	cdc.RegisterConcrete(&MsgAddWebsiteToProfile{}, "profile/AddWebsiteToProfile", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProfile{},
		&MsgUpdateProfile{},
		&MsgDeleteProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddBioToProfile{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddWebsiteToProfile{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
