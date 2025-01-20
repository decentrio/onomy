# sleep 7
# onomyd tx oracle set-price NOM 0.03 --home=$HOME/.onomyd/validator3  --from validator3 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

onomyd tx vaults create-vault 15000000000000000000000anom 50000000fxEUR --from test2 --home=$HOME/.onomyd/validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

sleep 7

onomyd q bank balances $(onomyd keys show test2 -a --keyring-backend test --home $HOME/.onomyd/validator1)

onomyd tx oracle set-price NOM 0.02 --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 20stake --chain-id onomy-mainnet-1 -y

