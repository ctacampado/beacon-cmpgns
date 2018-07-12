package main

import (
	"encoding/json"
	"log"
	"time"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/satori/go.uuid"
)

func addCampaign(fargs CCFuncArgs) pb.Response {
	log.Printf("starting addCampaign\n")

	u := uuid.Must(uuid.NewV4())
	var campaignID = u.String()

	c := CampaignInfo{CampaignID: campaignID, DateCreated: string(time.Now().Format("2006-Jan-02"))}

	err := json.Unmarshal([]byte(fargs.msg.Params), &c)
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
	log.Printf("bytes: %+v\n", bytes)
	err = fargs.stub.PutState(c.CampaignID, bytes)
	if err != nil {
		log.Printf("[addCampaign] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}
	fargs.msg.Data = string(bytes)
	log.Printf("fargs: %+v\n", fargs.msg)
	rspbytes, err := json.Marshal(fargs.msg)
	if err != nil {
		log.Printf("[addCampaign] Could not marshal fargs object: %+v\n", err)
		return shim.Error(err.Error())
	}
	log.Println("- end addCampaign")
	fargs.stub.SetEvent("newcmpgns", rspbytes)
	log.Printf("rspbytes: %+v\n", rspbytes)
	return shim.Success(rspbytes) //change nil to appropriate response
}
