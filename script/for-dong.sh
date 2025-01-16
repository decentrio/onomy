rm -rf $HOME/.onomyd
cp -r $HOME/1/.onomyd $HOME/

onomyd start --home=$HOME/.onomyd/validator1 &
onomyd start --home=$HOME/.onomyd/validator2 &
onomyd start --home=$HOME/.onomyd/validator3
