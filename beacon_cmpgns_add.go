package main

import (
	"encoding/json"
	"fmt"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/satori/go.uuid"
)

func addCampaign(fargs CCFuncArgs) pb.Response {
	log.Printf("starting addCampaign\n")

	u := uuid.Must(uuid.NewV4())
	var campaignID = u.String()

	c := CampaignInfo{CampaignID: campaignID}

	err := json.Unmarshal([]byte(fargs.req.Params), &c)
	if err != nil {
		return shim.Error("[addCampaign] Error unable to unmarshall msg: " + err.Error())
	}

	c.Status = "PLEDGE"
	c.DonatedAmount = "0"
	c.DisbursedAmount = "0"
	c.CampCompDate = "-"
	c.RatingFive = 0
	c.RatingFour = 0
	c.RatingThree = 0
	c.RatingTwo = 0
	c.RatingOne = 0

	log.Printf("[addCampaign ] campaign info: %+v\n", c)

	bytes, err := json.Marshal(c)
	if err != nil {
		log.Printf("[addCampaign] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(c.CampaignID, bytes)
	if err != nil {
		log.Printf("[addCampaign] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	fmt.Println("- end addCampaign")
	return shim.Success(nil) //change nil to appropriate response
}
