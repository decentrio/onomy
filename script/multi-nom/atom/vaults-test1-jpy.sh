sleep 7
onomyd tx oracle set-price ATOM 8.0 --home=$HOME/.onomyd/validator3  --from validator3 --keyring-backend test --fees 20stake --chain-id testing-1 -y

onomyd tx vaults create-vault 12500000atom 7575000000fxJPY --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y
sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price ATOM 5 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

sleep 31
onomyd q bank balances $(onomyd keys show test1 --home=$HOME/.onomyd/validator1  --keyring-backend test -a)

onomyd tx auction bid 0 2272000000fxJPY 908.8 --from test1 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

sleep 7

echo "I will buy all the remaining"
onomyd tx auction bid 0 9088000000fxJPY 753.3 --from test3 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id testing-1 -y

echo "wating long time, query auction ratecurrent = 1.1...liquidate"
# onomyd tx aution 