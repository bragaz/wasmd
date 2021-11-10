package v020

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateStore performs in-place store migrations from v1.0 to v2.0 The
// migration includes:
//
//
func MigrateStore(
	ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec,
) error {
	return nil
}
