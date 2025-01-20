onomyd tx vaults create-vault 12500000ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9 50000000fxEUR --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

# onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price ATOM 4 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 31
# onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

# onomyd tx auction bid 0 15000000fxEUR 1.2 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

# sleep 7

echo "I will buy all the remaining"
onomyd tx auction bid 0 47500000fxEUR 0.95 --from test3 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

echo "wating long t
ime, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 