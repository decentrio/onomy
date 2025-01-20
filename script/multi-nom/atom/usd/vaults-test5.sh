onomyd tx oracle set-price ATOM 8.0 --home=$HOME/.onomyd/validator3  --from validator3 --keyring-backend test --fees 20stake --chain-id testing-1 -y

onomyd tx vaults create-vault 12500000atom 50000000fxUSD --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y
# onomyd tx vaults create-vault 10000000atom 20000000fxUSD --from validator2 --home=$HOME/.onomyd/validator2 --keyring-backend test --fees 20stake --chain-id testing-1 -y

# sleep 7

# onomyd tx vaults mint 0 20000000fxUSD --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

# sleep 7 

# onomyd tx vaults deposit 0 20000000atom --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

# sleep 7

# onomyd tx vaults withdraw 0 1000atom --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

# sleep 7

# onomyd tx vaults repay 0 40000000fxUSD --from validator1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price ATOM 4 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

sleep 38
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 47500000fxUSD 0.95 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 