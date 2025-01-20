onomyd tx oracle set-price NOM 0.03 --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y
sleep 7
onomyd tx vaults create-vault 15000000000000000000000anom 50000000fxUSD --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y
sleep 7
onomyd tx oracle set-price NOM 0.1 --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y
sleep 7
onomyd tx vaults create-vault 5000000000000000000000anom 50000000fxUSD --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price NOM 0.02 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 8000000fxUSD 1 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 1 8000000fxUSD 1 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 