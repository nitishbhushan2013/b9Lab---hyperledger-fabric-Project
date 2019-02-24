# b9Lab---hyperledger-fabric-Project

The final project is an exercise designed to highlight your understanding of the whole technology stack and your command of all the different tools.

The Final Exam is split up into 4 parts.

1. Create an Use Case Art forgery is a major problem for the art market:
- Explain the problem and mention current solutions.
- Describe the advantages and disadvantages of using a blockchain solution.
- Design a network to help the art market. Think about Peers, CA, Channels and Smart Contracts. Make a critical analysis. Try to put all this on less than one page, usecase.md or usecase.txt.

2. Create a Network Assume, we have two organisations Org1 and Org2. Each organisation runs two peers in the network.
- Create certificates using cryptogen.
- Create a basic network using Docker Compose.
- Add a channel testchannel for Org1 and Org2 to your network.
- Write a chaincode for a simple asset transfer and deploy it.
- Use Fabric CA to replace the certificates.
- Enable TLS.
- Please give your asset at least two properties and use JSON to serialise it in the blockchain.
- Write a query function to prove asset transfers. Include comments in your source code.
Create a folder for this project. Upload everything you use and everything you generate.

3. Extend the Network Add Org3 with one peer to your network. You can set TLS disabled in your network.
- Use configtxlator to update the channel configuration. How does it work, how can you join Org3 to testchannel?
- Write a client for Org3 using NodeJS to invoke your chaincode.
- Org3 may also use the client to join its peer into the channel.
- Use Kafka OS.
- For this part, create a new folder. Upload everything you use and everything you generate. Put your explanation for configtxlator in a file called configtxlator.md or configtxlator.txt.

4. Decentralisation If you want to run the network on multiple hosts, you have a few options.
- You could run the nodes native on hosts. Why is this impractical?
- You could use various services, e.g. Docker Swarm or Kubernetes. Explain briefly Swarm and Kubernetes.
- Name some sources for the next steps to use them for Hyperledger Fabric.
- You can use IBM Cloud. Tell us what it offers in regard to Hyperledger Fabric.
- There is a project called Hyperledger Cello. Explain, what it aims to do and where one can learn more about it, as well how to get support.
- Try to put all this on less than one page, decentralisation.md or decentralisation.txt.
