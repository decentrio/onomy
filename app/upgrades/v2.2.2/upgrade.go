// Package v1_1_4 is contains chain upgrade of the corresponding version.
package v2_1_1 //nolint:revive,stylecheck // app version

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/onomyprotocol/onomy/app/keepers"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

// Name is migration name.
const Name = "v2.2.2"

// UpgradeHandler is an x/upgrade handler.
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {

	return func(c context.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		vm, err := mm.RunMigrations(ctx, configurator, vm)
		if err != nil {
			return vm, err
		}

		keepers.VaultsKeeper.Vaults.Walk(ctx, nil, func(key uint64, value vaultstypes.Vault) (stop bool, err error) {
			value.CreateTime = ctx.BlockTime()
			keepers.VaultsKeeper.Vaults.Set(ctx, key, value)
			return false, nil
		})

		return vm, nil
	}
}
