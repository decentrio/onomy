sleep 7
onomyd tx oracle set-price NOM 0.03 --home=$HOME/.onomyd/validator2  --from validator2 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y
onomyd tx vaults create-vault 125000000000000000000anom 50000000fxJPY --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 30stake --chain-id onomy-mainnet-1 -y --gas 300000

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price NOM 0.015 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 50000000fxJPY 1 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 