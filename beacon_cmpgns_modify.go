package main

import (
	"encoding/json"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func modifyCampaign(fargs CCFuncArgs) pb.Response {
	log.Printf("starting modifyIdentity\n")

	//get identity to be modified
	var qparams = &CampaignParams{}
	err := json.Unmarshal([]byte(fargs.req.Params), qparams)
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
	qr := getCOCampaigns(CCFuncArgs{stub: fargs.stub, req: Message{Params: string(qpbytes)}})
	err = json.Unmarshal([]byte(qr.Payload), qresp)
	if err != nil {
		return shim.Error("[modifyCampaign] Error unable to unmarshall msg: " + err.Error())
	}

	var campaign = &CampaignInfo{}
	err = json.Unmarshal([]byte(qresp.Elem[0].Value), campaign)
	if err != nil {
		return shim.Error("[modifyCampaign] Error unable to unmarshall msg: " + err.Error())
	}

	applyIdentityModsFromParam(qparams, campaign)

	cbytes, err := json.Marshal(*campaign)
	if err != nil {
		log.Printf("[modifyCampaign] Could not marshal campaign info object: %+v\n", err)
		return shim.Error(err.Error())
	}

	err = fargs.stub.PutState(campaign.CampaignID, cbytes)
	if err != nil {
		log.Printf("[modifyCampaign] Error storing data in the ledger %+v\n", err)
		return shim.Error(err.Error())
	}

	return shim.Success(nil) //change nil to appropriate response
}
