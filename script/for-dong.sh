# rm -rf $HOME/.onomy
# cp -r $HOME/4/.onomy $HOME/

# sleep 2
# onomyd start

###############################################################

rm -rf $HOME/.onomyd
cp -r $HOME/4/.onomyd $HOME/

sleep 2
onomyd start --home=$HOME/.onomyd/validator1 &
onomyd start --home=$HOME/.onomyd/validator2 &
onomyd start --home=$HOME/.onomyd/validator3
