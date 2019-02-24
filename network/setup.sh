!/bin/bash +x
#
# This fill will be used to quickly generate genesis block, create block configuration
# tx file and anchor peer update.
# 

#Create orderer genesis block
configtxgen -profile hlf11OrdererGenesis -outputBlock ./orderer/genesis.block

#Create channel configuration transaction file for each of the channel profile defined in the yaml file.
configtxgen -profile testChannel -outputCreateChannelTx ./channels/testchannel.tx -channelID testchannel

# channel - testchannel
#participating  orgs - org1 and org2
#Aim - Create one anchor peer for each org in this channel.
#Need to execute two tx block each with org1MSP and org2MSP respectively. 
configtxgen -profile testChannel -outputAnchorPeersUpdate ./channels/org1_anchor.tx -channelID testchannel -asOrg org1MSP
configtxgen -profile testChannel -outputAnchorPeersUpdate ./channels/org2_anchor.tx -channelID testchannel -asOrg org2MSP
