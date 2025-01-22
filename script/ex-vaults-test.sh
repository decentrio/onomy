onomyd tx oracle set-price NOM 0.03 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7
onomyd tx vaults create-vault 1250000000000000000000000anom 500000000000fxJPY --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 30stake --chain-id onomy-mainnet-1 -y --gas 300000

sleep 7

onomyd tx oracle set-price NOM 0.015 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31

onomyd tx auction bid 0 5000000000000fxJPY 1.4 --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 