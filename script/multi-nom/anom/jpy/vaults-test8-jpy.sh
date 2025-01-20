onomyd tx vaults create-vault 15000000000000000000000anom 7575000000fxJPY --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 30stake --chain-id onomy-mainnet-1 -y --gas 300000

sleep 7
onomyd tx oracle set-price NOM 0.1 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

onomyd tx vaults create-vault 5000000000000000000000anom 7575000000fxJPY --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 30stake --chain-id onomy-mainnet-1 -y --gas 300000

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price NOM 0.02 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 1212000000fxJPY 1 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

echo "I will buy all the remaining"
onomyd tx auction bid 1 1212000000fxJPY 1 --from test3 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

echo "wating long time, query auction ratecurrent = 1.1...liquidate"