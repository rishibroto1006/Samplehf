package main

import (
	"fmt"
	"proj1/blockchain"
	"proj1/web"
	"proj1/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.hf.Rishi.com",

		// Channel parameters
		ChannelID:     "channelone",
		ChannelConfig: os.Getenv("GOPATH") + "/src/proj1/fixtures/artifacts/channelone.tx",

		// Chaincode parameters
		ChainCodeID:     "hf-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "proj1/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}
