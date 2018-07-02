package main

import (
	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//--------------------------------------------------------------------------
//Start adding Chaincode-related Structures here

//CCFuncArgs common cc func args
type CCFuncArgs struct {
	function string
	req      Message
	stub     shim.ChaincodeStubInterface
}

type ccfunc func(args CCFuncArgs) pb.Response

//Chaincode cc structure
type Chaincode struct {
	FMap map[string]ccfunc //ccfunc map
	Msg  Message           //data
}

//COCCMessage Charity Org Chain Code Message Structure
type Message struct {
	CID    string `json:"CID"`    //ClientID --for websocket push (event-based messaging readyness)
	AID    string `json:"AID"`    //ActorID (Donor ID/Charity Org ID/Auditor ID/etc.)
	Type   string `json:"type"`   //Chaincode Function
	Params string `json:"params"` //Function Parameters
}

//End of Chaincode-related Structures
//--------------------------------------------------------------------------
//Start adding Query Parameter (Parm) Structures here

type QRes struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type QRsp struct {
	Elem []QRes `json:"elem"`
}

//CampaignQueryParams Structure for Query Parameters
type CampaignParams struct {
	CharityID       string `json:"CharityID,omitempty"`
	CampaignID      string `json:"CampaignID,omitempty"`
	Status          int    `json:"Status,omitempty"`
	CampStartDate   string `json:"CampStartDate,omitempty"`
	DonatedAmount   string `json:"DonatedAmount, omitempty"`
	DisbursedAmount string `json:"DisbursedAmount, omitempty"`
	RatingFive      int    `json:"RatingFive, omitempty"`
	RatingFour      int    `json:"RatingFour, omitempty"`
	RatingThree     int    `json:"RatingThree, omitempty"`
	RatingTwo       int    `json:"RatingTwo, omitempty"`
	RatingOne       int    `json:"RatingOne, omitempty"`
}

//COCCQuerySelector Structure for Query Selector
type CampaignQuerySelector struct {
	Selector CampaignParams `json:"selector"`
}

//End of Query Paramter Structures
//--------------------------------------------------------------------------
//Start adding Data Models here

//CampStatus type for Enum
type CampStatus int

//Campaign Status Enum
const (
	PLEDGE CampStatus = 1 + iota
	DISBURSE
	COMPLETED
	NEW
	CANCELED
)

//CampaignInfo data model
type CampaignInfo struct {
	CampaignID      string     `json:"CampaignID, omitempty"`
	CharityName     string     `json:"CharityName, omitempty"`
	CharityID       string     `json:"CharityID"`
	CampaignName    string     `json:"CampaignName"`
	Description     string     `json:"Description"`
	CampaignCaption string     `json:"CampaignCaption, omitempty"`
	CampStartDate   string     `json:"CampStartDate"`
	CampEndDate     string     `json:"CampEndDate"`
	CampCompDate    string     `json:"CampCompDate, omitempty"`
	CampaignPhoto   string     `json:"CampaignPhoto, omitempty"`
	Status          CampStatus `json:"Status"`
	CampaignAmount  string     `json:"CampaignAmount"`
	DonatedAmount   string     `json:"DonatedAmount, omitempty"`
	DisbursedAmount string     `json:"DisbursedAmount, omitempty"`
	RatingFive      int        `json:"RatingFive, omitempty"`
	RatingFour      int        `json:"RatingFour, omitempty"`
	RatingThree     int        `json:"RatingThree, omitempty"`
	RatingTwo       int        `json:"RatingTwo, omitempty"`
	RatingOne       int        `json:"RatingOne, omitempty"`
}

//End of Data Models
//--------------------------------------------------------------------------