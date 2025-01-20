sleep 7
# onomyd tx oracle set-price NOM 0.03 --home=$HOME/.onomyd/validator3  --from validator3 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

onomyd tx vaults create-vault 15000000000000000000000anom 50000000fxEUR --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price NOM 0.02 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 15000000fxEUR 1.2 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

echo "I will buy all the remaining"
onomyd tx auction bid 0 60000000fxEUR 1 --from test3 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 