package main

import (
	"encoding/json"
	"fmt"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func modifyCampaign(fargs CCFuncArgs) pb.Response {
	log.Printf("starting modifyCampaign\n")

	//get identity to be modified
	var qparams = &CampaignParams{}
	err := json.Unmarshal([]byte(fargs.msg.Params), qparams)
	if err != nil {
		return shim.Error("[modifyCampaign] Error unable to unmarshall Params: " + err.Error())
	}

	var qp = &CampaignParams{CampaignID: qparams.CampaignID}
	qpbytes, err := json.Marshal(*qp)
	if err != nil {
		log.Printf("[modifyCampaign] Could not marshal campaignParams: %+v\n", err)
		return shim.Error(err.Error())
	}

	var qresp = &QRsp{}
	qr := getCOCampaigns(CCFuncArgs{stub: fargs.stub, msg: Message{Params: string(qpbytes)}})
	err = json.Unmarshal([]byte(qr.Payload), qresp)
	if err != nil {
		return shim.Error("[modifyCampaign] Error unable to unmarshall msg: " + err.Error())
	}

	var campaign = CampaignInfo{}
	err = json.Unmarshal([]byte(qresp.Elem[0].Value), &campaign)
	if err != nil {
		return shim.Error("[modifyCampaign] Error unable to unmarshall msg: " + err.Error())
	}

	applyIdentityModsFromParam(qparams, &campaign)
	log.Printf("campaign: %+v\n", campaign)
	cbytes, err := json.Marshal(campaign)
	if err != nil {
		log.Printf("[modifyCampaign] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(campaign.CampaignID, cbytes)
	if err != nil {
		log.Printf("[modifyCampaign] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	fargs.msg.Data = string(cbytes)
	rspbytes, err := json.Marshal(fargs.msg)
	fmt.Printf("- end modifyCampaign")
	fargs.stub.SetEvent("modcmpgns",rspbytes)
	return shim.Success(rspbytes) //change nil to appropriate response
}
