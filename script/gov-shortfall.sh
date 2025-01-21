onomyd tx gov submit-proposal ./script/proposal-3.json --home=$HOME/.onomyd/validator1  --from validator1 --keyring-backend test --fees 50stake --chain-id onomy-mainnet-1 -y --gas 500000

sleep 7

onomyd tx gov vote 1 yes  --from validator1 --keyring-backend test --home ~/.onomyd/validator1 --chain-id onomy-mainnet-1 -y --fees 20stake
onomyd tx gov vote 1 yes  --from validator2 --keyring-backend test --home ~/.onomyd/validator2 --chain-id onomy-mainnet-1 -y --fees 20stake
onomyd tx gov vote 1 yes  --from validator3 --keyring-backend test --home ~/.onomyd/validator3 --chain-id onomy-mainnet-1 -y --fees 20stake