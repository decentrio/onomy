// Package v1_1_4 is contains chain upgrade of the corresponding version.
package v2_2_4 //nolint:revive,stylecheck // app version

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/onomyprotocol/onomy/app/keepers"
)

// Name is migration name.
const Name = "v2.2.4"

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

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "ATOM", "USD", 6, 6)
		if err != nil {
			return vm, err
		}

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "ATOM", "EUR", 6, 6)
		if err != nil {
			return vm, err
		}

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "ATOM", "JPY", 6, 6)
		if err != nil {
			return vm, err
		}

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "NOM", "USD", 18, 6)
		if err != nil {
			return vm, err
		}

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "NOM", "EUR", 18, 6)
		if err != nil {
			return vm, err
		}

		err = keepers.OracleKeeper.SetPairDecimalsRate(c, "NOM", "JPY", 18, 6)
		if err != nil {
			return vm, err
		}

		return vm, nil
	}
}
