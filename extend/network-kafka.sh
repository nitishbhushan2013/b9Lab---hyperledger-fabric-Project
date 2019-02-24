!/bin/bash +x
#
# This fill will does below activities
# 1. channel cteation
# 2. peer join the channel
# 3. peer channel update for anchor peer 
# 4. chaincode installation
#5. chaoincode initialization

#create channel from org1 cli  
read -p "press any key to create channel"

docker exec cli.org1 bash -c 'peer channel create -c testchannel -f ./channels/testchannel.tx -o orderer2.hlf11.com.au:7050'



read -p "press any key to let peers join the testchannel"
#let peer join three channels
docker exec cli.org1 bash -c 'peer channel join -b testchannel.block'
docker exec cli.org2 bash -c 'peer channel join -b testchannel.block'


read -p "press any key to let individual org update its anchor peer in the respective channel"
#let org updates its anchor peer in the channel
docker exec cli.org1 bash -c 'peer channel update -o orderer2.hlf11.com.au:7050 -c testchannel -f ./channels/Org1MSPanchors.tx'
docker exec cli.org2 bash -c 'peer channel update -o orderer2.hlf11.com.au:7050 -c testchannel -f ./channels/Org2MSPanchors.tx'

read -p "press any key to install chaincode on each of the peer"
#let install chanincode on each peer
docker exec cli.org1 bash -c 'peer chaincode install -p artWork -n artWork -v 0'
docker exec cli.org2 bash -c 'peer chaincode install -p artWork -n artWork -v 0'

read -p "press any key to instantiate chaincode on each of the peer"
#let instantiate chaincode on each peer
docker exec cli.org1 bash -c "peer chaincode instantiate -C testchannel -n artWork -v 0 -c '{\"Args\":[]}’"
docker exec cli.org2 bash -c "peer chaincode instantiate -C testchannel -n artWork -v 0 -c '{\"Args\":[]}’"


