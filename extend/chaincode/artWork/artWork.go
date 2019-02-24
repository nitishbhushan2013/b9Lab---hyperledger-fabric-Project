package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim" // import for Chaincode Interface
	pb "github.com/hyperledger/fabric/protos/peer"      // import for peer response
)

//ArtWork Defined to implement chaincode interface
type ArtWork struct {
}

//ArtItem : An Asset represent the datastructure for the Asset object.
type ArtItem struct {
	ArtID    string // unique Id for artwork
	ArtName  string // art name
	ArtURL   string // art url
	ArtStyle string // art style - 'watermark', 'oil', 'sketch'
	Artist   string // Artist name
	Owner    string // current Owner name

}

// Init: Implement Init
func (c *ArtWork) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke: Implement Invoke
func (c *ArtWork) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters() // get function name and args

	switch function {
	case "createArtWork":
		// create the art work
		return c.createArtWork(stub, args)
	case "transferArtWork":
		// transfer the art work to other owner
		return c.transferArtWork(stub, args)
	case "queryArtWork":
		// query the artwork by artID
		return c.queryArtWorkByArtID(stub, args)
	case "listByOwner":
		// get list of all the artwork for the given onwner
		return c.listByOwner(stub, args)
	default:
		return shim.Error("Available functions: createArtWork, transferArtWork, queryArtWork, listByOwner")
	}
}

// createArtWork will create the artwork based on the given arguments
func (c *ArtWork) createArtWork(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 6 {
		return shim.Error("createArtWork arguments usage: ArtID, ArtName, ArtURL, ArtStyle, Artist, Owner")
	}

	// new artwork is created
	ArtItemObj := ArtItem{ArtID: args[0], ArtName: args[1], ArtURL: args[2], ArtStyle: args[3], Artist: args[4], Owner: args[5]}

	// Use JSON to store in the Blockchain
	artItemAsBytes, err := json.Marshal(ArtItemObj)

	if err != nil {
		return shim.Error(err.Error())
	}

	// Use ArtID as key
	err = stub.PutState(ArtItemObj.ArtID, artItemAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// transferArtWork will transfer the given artwork from current owner to new owner
func (c *ArtWork) transferArtWork(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("transferArtWork arguments usage: ArtID, Owner")
	}

	// get ArtID
	artItemAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ArtID " + args[0] + " not found ")
	}

	// Get Information from Blockchain
	artItem := ArtItem{}
	// Decode JSON data
	json.Unmarshal(artItemAsBytes, &artItem)

	// Change the owner
	artItem.Owner = args[1]
	// Encode JSON data
	artItemAsBytes, _ = json.Marshal(artItem)
	// Store in the Blockchain
	err = stub.PutState(artItem.ArtID, artItemAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

// queryArtWorkByArtID will give back the artwork for the given artId
func (c *ArtWork) queryArtWorkByArtID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("queryArtWorkByArtID arguments usage: ArtId")
	}

	// get ArtId
	artItemAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ArtId " + args[0] + " not found ")
	}

	return shim.Success(artItemAsBytes)

}

// listByOwner would list all the art works belongs to given owner
func (c *ArtWork) listByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("listByOwner arguments usage: Owner")
	}

	// get ArtId
	artItemAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("ArtId " + args[0] + " not found ")
	}

	return shim.Success(artItemAsBytes)
}

func main() {
	err := shim.Start(new(ArtWork))
	if err != nil {
		fmt.Printf("Error starting chaincode sample: %s", err)
	}
}
