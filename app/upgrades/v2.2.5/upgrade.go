// Package v1_1_4 is contains chain upgrade of the corresponding version.
package v2_2_5 //nolint:revive,stylecheck // app version

import (
	"context"
	"slices"
	"strings"

	"cosmossdk.io/collections"

	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/onomyprotocol/onomy/app/keepers"
	auctiontypes "github.com/onomyprotocol/reserve/x/auction/types"
	vaultstypes "github.com/onomyprotocol/reserve/x/vaults/types"
)

// Name is migration name.
const Name = "v2.2.5-testnet"

var (
	fxUSD = "fxUSD"
	fxEUR = "fxEUR"
	fxJPY = "fxJPY"
)

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

		// genesis

		// bank nom -> fx
		err = keepers.BankKeeper.Balances.Walk(ctx, nil, func(key collections.Pair[sdk.AccAddress, string], value math.Int) (stop bool, err error) {
			addr := sdk.AccAddress(key.K1())
			if key.K2() == "nomUSD" {
				err = keepers.BankKeeper.Balances.Set(ctx, collections.Join(addr, fxUSD), value)
				if err != nil {
					return true, err
				}

				err = keepers.BankKeeper.Balances.Remove(ctx, key)
				if err != nil {
					return true, err
				}
			}
			if key.K2() == "nomEUR" {
				err = keepers.BankKeeper.Balances.Set(ctx, collections.Join(addr, fxEUR), value)
				if err != nil {
					return true, err
				}

				err = keepers.BankKeeper.Balances.Remove(ctx, key)
				if err != nil {
					return true, err
				}
			}
			if key.K2() == "nomJPY" {
				err = keepers.BankKeeper.Balances.Set(ctx, collections.Join(addr, fxJPY), value)
				if err != nil {
					return true, err
				}

				err = keepers.BankKeeper.Balances.Remove(ctx, key)
				if err != nil {
					return true, err
				}
			}

			return false, nil
		})
		if err != nil {
			return vm, err
		}

		// vault
		err = keepers.VaultsKeeper.SetParams(ctx, vaultstypes.DefaultParams())
		if err != nil {
			return vm, err
		}

		err = keepers.VaultsKeeper.VaultsManager.Walk(ctx, nil, func(key string, value vaultstypes.VaultManager) (stop bool, err error) {
			if !slices.Contains(vaultstypes.DefaultMintDenoms, value.Params.MintDenom) {
				value.Params.MintDenom = paseMintDenom(value.Params.MintDenom)
				err = keepers.VaultsKeeper.VaultsManager.Set(ctx, key, value)
				if err != nil {
					return true, err
				}
			}
			return false, nil
		})
		if err != nil {
			return vm, err
		}

		err = keepers.VaultsKeeper.Vaults.Walk(ctx, nil, func(key uint64, value vaultstypes.Vault) (stop bool, err error) {
			if !slices.Contains(vaultstypes.DefaultMintDenoms, value.Debt.Denom) {
				value.Debt.Denom = paseMintDenom(value.Debt.Denom)
				err = keepers.VaultsKeeper.Vaults.Set(ctx, key, value)
				if err != nil {
					return true, err
				}
			}
			return false, nil
		})
		if err != nil {
			return vm, err
		}
		// auction
		keepers.AuctionKeeper.Auctions.Walk(ctx, nil, func(key uint64, value auctiontypes.Auction) (stop bool, err error) {
			if !slices.Contains(vaultstypes.DefaultMintDenoms, value.TokenRaised.Denom) {
				value.TokenRaised.Denom = paseMintDenom(value.TokenRaised.Denom)
				value.TargetGoal.Denom = paseMintDenom(value.TargetGoal.Denom)
				err = keepers.AuctionKeeper.Auctions.Set(ctx, key, value)
				if err != nil {
					return true, err
				}
			}

			return false, nil
		})

		keepers.AuctionKeeper.Bids.Walk(ctx, nil, func(key uint64, value auctiontypes.BidQueue) (stop bool, err error) {
			for i, bid := range value.Bids {
				if !slices.Contains(vaultstypes.DefaultMintDenoms, bid.Amount.Denom) {
					bid.Amount.Denom = paseMintDenom(bid.Amount.Denom)
					value.Bids[i] = bid
				}
			}

			return false, nil
		})

		return vm, nil
	}
}

func paseMintDenom(denom string) string {
	if strings.Contains(denom, "JPY") {
		return fxJPY
	}
	if strings.Contains(denom, "EUR") {
		return fxEUR
	}
	return fxUSD
}
