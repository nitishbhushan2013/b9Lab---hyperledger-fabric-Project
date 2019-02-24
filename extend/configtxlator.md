Hyperledger fabric current state - 
--------------------------------
Hyper ledger fabric is running with defined organization and its peers. Channel has been created to contains the peers from all the participating organisations.

Current requirnment - 
--------------------
There is a requirnment to add a new organisation - org3, and its peer into already running hyperldger fabric network. 

We can use configtxlator tool to achieve the above objective.

configtxlator - 
--------------
Configtxlaror tool helps us to update channel configuration or genesis block. So it can be used to add new organization to the channel. 

It allows users to translate between protobuf and JSON versions of fabric data structures and create config updates. The command may either start a REST server to expose its functions over HTTP or may be utilized directly as a command line tool.
The tool converts easily between different equivalent data representations/formats. For example, in one mode of tool operation, the tool performs conversions between the binary protobuf format to a human-readable JSON textual format, and vice-versa. Additionally, the tool can compute configuration updates based on the differences between two different sets of configurations transactions.

Below are the standard usages of it 
    - SDK retrieves latest config 
    - configtxlator produces human readable JSON version of config
    - On need basis, user update the config details
    - configtxlator compute the delta change by comparing current config with the original config
    - SDK submit signs and submit config


Technical prespective : https://hyperledger-fabric.readthedocs.io/en/release-1.2/commands/configtxlator.html
----------------------

The configtxlator command allow user to translate between protobuf and JSON version of fabric data structure and compute config updates. The command starts a REST server to expose its functionalities over HTTP or can be invoked directly over command line interface. 

It has below subcomponents 
    - start : start the REST server over http
    - proto_encode : comnvert the JSON document to protobuf
    - proto_decode : convert the protobuf to JSON document
    - compute_update : take two marshalled config messages and compute the config upadte between these two 
    - version : show version information 



----------------------------------------------------------------------------------------------------------
              Steps of adding a new organisation into fabric runtime using configtxlator
----------------------------------------------------------------------------------------------------------
Important links 
https://developer.ibm.com/tutorials/cl-add-an-organization-to-your-hyperledger-fabric-blockchain/
https://hyperledger-fabric.readthedocs.io/en/v1.1.0-alpha/channel_update.html
https://www.youtube.com/watch?v=DKuGU5CYV_E


Minimum req -
-----------
Version 1.1.0-preview of Hyperledger Fabric or higher

Supporting tool -
---------------
jq tool - open source jq tool to script interactions with the JSON returned by configtxlator.

High level steps 
-----------------
** Setting up env **
step 0: Start the configtxlator tool in the background, and verify that the tool has started correctly to receive incoming client requests


** House Keeping **
Step1: create org3 folder

Step2 : Create  file in it - crypto-config.yaml and configtx.yaml for the org3
Step3: Generate crypto materials for Org3
Step4: Get the org3MSP definition in JSON format : Now use the the configtxgen utility to print out the Org3-specific configuration material in JSON representation.
Step5: The final piece of housekeping activity - port the Orderer Org’s MSP material into the Org3 crypto-config directory. 



** Steps of adding a new organisation **
Step6: Retrieve the current configuration.
step7: Decode the configuration into a human-readable version of JSON configuration using configtxlator.
step8: Extract the config section.
step9: Create the new configuration by performing automated or manual edits on the extracted config section.
step10: Encode both updated and original configurations using configtxlator.
step11: Send them to configtxlator to compute the config update delta, which represents the changes to the config.
step12: Decode the config update and wrap it up into a config update envelope.
step13: Create the new config transaction.
step14 : Update the channel by submitting the new signed config transaction.
Step15 : Export the Org2 environment variables:
Step16: And lastly we will issue the peer channel update command. The Org2 Admin signature will be attached to this call, so there is no need to manually sign the proto a second time:
step17: JOIN ORG3 TO THE CHAHHEL 
Step17.1 : Get the genesis block 
Step18 : Join the genesis block 
Step 19 Upgrade & invoke chaincode
step 19.1 : Install the new update of the chaincode on org3
Step19.2 : Now jump back to the original CLI container and install the new version on the Org1 and Org2 peers.
Step 19.3 : Export the Org2 environment variables:
Step 19.4 : install the new version on the org2 peer
Step20: Upgrade the chaincode to include org3 to the emdorsement policy

-------------------------------------------------------------------------------------------------------------------
                        ----------------------------------
                        DETAILED EXPLANATION for EACH STEPS
                        ----------------------------------
-------------------------------------------------------------------------------------------------------------------



Setting up env
---------------
STEP 0: Start the configtxlator tool in the background, and verify that the tool has started correctly to receive incoming client requests

exec into the CLI container.  The bootstrapped identity is the Org1 admin user, meaning that any steps where we want to act on behalf of Org2 will require the export of MSP-specific environment variables.

check jq tool -  apt update && apt install jq

Then, docker exec -it cli.org1  bash

### Nitishs-MacBook-Pro:hlf11 nitishbhushan$ docker exec -it cli.org1 bash
root@a42f7983ea4a:/opt/gopath/src/github.com/hyperledger/fabric/peer#

then execute below command to start configtxlator as REST service
root@a42f7983ea4a:/opt/gopath/src/github.com/hyperledger/fabric/peer# configtxlator start &


Step1: create org3 folder
-----------------------
Step2 : Create  file in it - crypto-config.yaml and configtx.yaml for the org3
-------------------------------------------------------------------------------
Step3: Generate crypto materials for Org3
-----------------------------------------
go to /org3 folder

Nitishs-MacBook-Pro:org3 nitishbhushan$ cryptogen generate --config crypto-config.yaml
org3.com.au


Step4: Get the org3MSP definition in JSON format : Now use the the configtxgen utility to print out the Org3-specific configuration material in JSON representation.
------------------------------------------------------------------------------
configtxgen -printOrg org3MSP >> ../channel-artifacts/org3.json

The above command creates a JSON file - org3.json - and outputs it into the channel-artifacts subdirectory at the root of first-network. This file contains the modification policy definitions for Org3, as well as three important certificates presented in base 64 format: admin user cert, CA root cert and TLS root cert.


Step5: The final piece of housekeping activity - port the Orderer Org’s MSP material into the Org3 crypto-config directory.
------------------------------------------------------------------------------
In particular, we are concerned with the Orderer’s TLS root cert, which will allow for secure communication between Org3 entities and the network’s ordering node.

cd .. && cp -r crypto-config/ordererOrganizations org3/crypto-config


Step6: Retrieve the current configuration.
------------------------------------------
Login to cli console 
docker exec -it cli.org1 bash and 
1. execute below command 
peer channel fetch config config_block.pb  -o orderer.hlf11.com.au:7050 -c testchannel 


When you issued the peer channel fetch command, there was a decent amount of output in the terminal. The last line in the logs is of interest:
2018-10-27 05:33:13.596 UTC [channelCmd] readBlock -> DEBU 011 Received block: 2

This is telling us that the most recent configuration block for mychannel is actually block 2, NOTthe genesis block. By default, the peer channel fetch config command returns the most recent configuration block for the targeted channel, which is block 2 in our case. When the fabric network got created, the embedded script made two additional configuration updates to the channel. Namely, anchor peers for our two organizations - Org1 & Org2 - were defined by means of two separate channel update transactions. As such, we have the following configuration sequence: block 0 - genesis; block 1 - Org1 anchor peer update; block 2 - Org2 anchor peer update.


step7: Decode the configuration into a human-readable version of JSON configuration using configtxlator.
---------------------------------------------------------------------
curl -X POST --data-binary @config_block.pb http://127.0.0.1:7059/protolator/decode/common.Block > config_block.json


step8: Extract the config section.
---------------------------------
Extract the config section of data's payload data section from the decoded channel configuration block for application channel mychannel, and verify the correct and successful extraction

jq .data.data[0].payload.data.config config_block.json > config.json


step9: Create the new configuration by performing automated or manual edits on the extracted config section.
----------------------------------------------------------------------------------
manual edit the config.json ( /Users/nitishbhushan/Documents/nitish_operations/technical/tehnologies/hyperledger-fabric/b9lab_fabric/certification/b9Lab-HLF11/extend/chaincode/github.com/hyperledger/fabric/peer/config.json) and add the org3.json details in it 

Or execute below command to do the same thing automatically using jq
1. copy channel-artifacts/org3.json to /Users/nitishbhushan/Documents/nitish_operations/technical/tehnologies/hyperledger-fabric/b9lab_fabric/certification/b9Lab-HLF11/extend/chaincode/github.com/hyperledger/fabric/peer

2. execute below command
jq -s '.[0] * {"channel_group":{"groups":{"Application":{"groups": {"Org3MSP":.[1]}}}}}' config.json org3.json >& updated_config.json


step10: Encode both updated and original configurations using configtxlator.
--------------------------------------------------------------------------
1. Set the URL:
CONFIGTXLATOR_URL=http://127.0.0.1:7059
export CHANNEL_NAME=testchannel
Export the ORDERER_CA and CHANNEL_NAME variables:

2. First, encode config.json to config.pb:
curl -X POST --data-binary @config.json "$CONFIGTXLATOR_URL/protolator/encode/common.Config" > config.pb


3. Next, encode updated_config.json to updated_config.pb:
curl -X POST --data-binary @updated_config.json "$CONFIGTXLATOR_URL/protolator/encode/common.Config" > updated_config.pb




step11: Send them to configtxlator to compute the config update delta, which represents the changes to the config.
-----------------------------------------------------------------------------------------------------------------
Compute the configuration update, which transitions between the original and modified channel configurations for the application channel mychannel using the tool, by executing the following command.

curl -X POST -F channel=$CHANNEL_NAME -F "original=@config.pb" -F "updated=@updated_config.pb" "${CONFIGTXLATOR_URL}/configtxlator/compute/update-from-configs" > org3_update.pb

This new proto - org3_update.pb - contains the Org3 definitions and high level pointers to the Org1 and Org2 material. We are able to forgo the extensive MSP material and modification policy information for Orgs 1 and 2, because this data is already present within the channel’s genesis block. As such, we only need the delta between the two configurations.




step12: Decode the config update and wrap it up into a config update envelope.
-------------------------------------------------------------------------------
Before submitting the channel update, we need to perform a few final steps. First, let’s decode this object into editable JSON format and call it org3_update.json:
curl -X POST --data-binary @org3_update.pb "$CONFIGTXLATOR_URL/protolator/decode/common.ConfigUpdate" | jq . > org3_update.json

Now, we have a decoded update file - org3_update.json - that we need to wrap in an envelope message. This step will give us back the header field that we stripped away earlier. We’ll name this file - org3_update_in_envelope.json:

echo '{"payload":{"header":{"channel_header":{"channel_id":"testchannel", "type":2}},"data":{"config_update":'$(cat org3_update.json)'}}}' | jq . > org3_update_in_envelope.json


step13: Create the new config transaction.
------------------------------------------


Using our properly formed JSON - org3_update_in_envelope.json - we will leverage the configtxlator tool one last time and convert this object into the fully fledged proto format that Fabric requires. We’ll name our final update object - org3_update_in_envelope.pb:


curl -X POST --data-binary @org3_update_in_envelope.json "$CONFIGTXLATOR_URL/protolator/encode/common.Envelope" > org3_update_in_envelope.pb


step14 : Update the channel by submitting the new signed config transaction.
----------------------------------------------------------------------------
We now have a protobuf binary - org3_update_in_envelope.pb - within our CLI container, however we need signatures from the requisite Admin users before we can successfully submit the update. The modification policy (mod_policy) for our channel is set to the default of “MAJORITY”, which means that we need an Admin from both of the initial organizations - Org1 & Org2 - to sign off on this update transaction. If we fail to obtain these two signatures, then the ordering service will reject the transaction for failing to fulfill the policy. First, let’s sign this update proto as the Org1 Admin. Remember that the CLI container is bootstrapped with the Org1 MSP material, so we simply need to issue the peer channel signconfigtx command:

peer channel signconfigtx -f org3_update_in_envelope.pb


===============================================================================================================

The final step is to switch the CLI container’s identity to reflect the Org2 Admin user. We do this by exporting four environment variables specific to the Org2 MSP.

Note: The following maneuver is not reflective of a real world operation. A single container would never be mounted with an entire network’s crypto material. Rather, the update object would need to be securely passed out-of-band to an Org2 Admin for inspection and approval.


Step15 : Export the Org2 environment variables:
----------------------------------------------
export CORE_PEER_LOCALMSPID="org2MSP"
export CORE_PEER_ADDRESS=peer0.org2.com.au:7051
export CHANNEL_NAME="testchannel"

export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com.au/peers/peer0.org2.com.au/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com.au/users/Admin@org2.com.au/msp

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/hlf11.com.au/orderers/orderer.hlf11.com.au/msp/tlscacerts/tlsca.hlf11.com.au-cert.pem


Step16: And lastly we will issue the peer channel update command. The Org2 Admin signature will be attached to this call, so there is no need to manually sign the proto a second time:
--------------------------------------------------------------------------------------------------------------
tls disabled:
---------------
peer channel update -f org3_update_in_envelope.pb -c $CHANNEL_NAME -o orderer.hlf11.com.au:7050 

tls enabled 
-------------
peer channel update -f org3_update_in_envelope.pb -c $CHANNEL_NAME -o orderer.example.com:7050 --tls --cafile $ORDERER_CA




step17: JOIN ORG3 TO THE CHAHHEL 
---------
: First, let’s launch the containers for the Org3 peers and an Org3-specific CLI.

docker-compose -f docker-compose-hlf11-org3.yaml up -d

Step 17.1
----------
Get the genesis block 

Now let’s send a call to the ordering service asking for the genesis block of testchannel. The ordering service is able to verify the signature attached to this call as a result of our successful channel update. If Org3 had not been successfully appended to the channel config, then the ordering service would reject this request.

tls disabled:
---------------
peer channel fetch 0 testchannel.block -o orderer.hlf11.com.au:7050 -c testchannel 

tls enabled:
---------------
peer channel fetch 0 testchannel.block -o orderer.hlf11.com.au:7050 -c testchannel --tls --cafile $ORDERER_CA


Note: Notice, that we are passing a 0 to indicate that we want the first block on the channel’s ledger (i.e. the genesis block). If we simply passed the peer channel fetch config command, then we would have received block 5 - the updated config with Org3 defined. However, we can’t begin our ledger with a downstream block; rather we need to join with block 0.


Step18 : Join the genesis block 
---------------------------------
Issue the peer channel join command and pass in the genesis block - testchannel.block:

peer channel join -b testchannel.block


Step 19 Upgrade & invoke chaincode
----------------------------------
step 19.1 : Install the new update of the chaincode on org3
-------------------------------------------------------------
The final piece of the puzzle is to increment the chaincode version and update the endorsement policy to include Org3. Stay in the Org3 CLI container and install the chaincode. 
Since we know that an upgrade is coming, we can forgo the futile exercise of installing version 0 of the chaincode. We are solely concerned with the new version where Org3 will be part of the endorsement policy, therefore we’ll jump directly to version 1:
peer chaincode install -p artWork -n artWork -v 1


Step19.2 : Now jump back to the original CLI container and install the new version on the Org1 and Org2 peers. 
--------------------------------------------------------------------------------------------------------------
docker exec -it cli.org1 bash
    peer chaincode install -p artWork -n artWork -v 1

Step 19.3 : Export the Org2 environment variables:
----------------------------------------------
export CORE_PEER_LOCALMSPID="org2MSP"
export CORE_PEER_ADDRESS=peer0.org2.com.au:7051
export CHANNEL_NAME="testchannel"

export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com.au/peers/peer0.org2.com.au/tls/ca.crt

export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.com.au/users/Admin@org2.com.au/msp

export ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/hlf11.com.au/orderers/orderer.hlf11.com.au/msp/tlscacerts/tlsca.hlf11.com.au-cert.pem

Step 19.4 : install the new version on the org2 peer
------------------------------------------------------
peer chaincode install -p artWork -n artWork -v 1


Step20: Upgrade the chaincode to include org3 to the emdorsement policy
There have been no modifications to the underlying source code, we are simply adding Org3 to the endorsement policy for a chaincode - artWork  on a channel -testchannel.


Execute below from org1, org2 and org3 cli console

tls disabled:
---------------
peer chaincode upgrade -o orderer.hlf11.com.au:7050 -n artWork -C testchannel -v 1 -c '{"Args":[]}’ -P "OR ('org1MSP.member','org2MSP.member','org3MSP.member')"

tls enabled:
---------------

peer chaincode upgrade -o orderer.hlf11.com.au:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C testchannel -n artWork -v 1 -c '{\"Args\":[]}’ -P "OR ('org1MSP.member','org2MSP.member','org3MSP.member')"


The upgrade call adds a new block to the channel’s ledger and allows for the Org3 peers to execute transactions during the endorsement phase. 


------------------------------------------------------------------------------------------------------------------
