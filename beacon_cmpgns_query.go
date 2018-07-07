package main

import (
	"encoding/json"
	"fmt"
	"log"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func getCOCampaigns(fargs CCFuncArgs) pb.Response {
	fmt.Println("starting getCOCampaigns")

	var qparams = &CampaignParams{}
	log.Printf("qparams start: %+v\n", qparams)
	err := json.Unmarshal([]byte(fargs.msg.Params), qparams)
	if err != nil {
		return shim.Error("[getCOCampaigns] Error unable to unmarshall msg: " + err.Error())
	}
	log.Printf("qparams after unmarshall: %+v\n", qparams)

	qstring, err := createQueryString(qparams)
	if err != nil {
		return shim.Error("[getCOCampaigns] Error unable to create query string: " + err.Error())
	}

	log.Printf("[getIdentity] query using querystring: %+v\n", qparams)
	resultsIterator, err := fargs.stub.GetQueryResult(qstring)
	fmt.Printf("- getQueryResultForQueryString resultsIterator:\n%+v\n", resultsIterator)
	defer resultsIterator.Close()
	if err != nil {
		return shim.Error("[getCOCampaigns] Error unable to GetQueryResult: " + err.Error())
	}

	var qresp = QRsp{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error("[getCOCampaigns] Error unable to get next item in iterator: " + err.Error())
		}

		q := QRes{Key: queryResponse.Key, Value: string(queryResponse.Value)}
		qresp.Elem = append(qresp.Elem, q)
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%+v\n", qresp)
	fmt.Printf("- getQueryResultForQueryString querystring:\n%s\n", qstring)
	fmt.Printf("- getQueryResultForQueryString qparams:\n%+v\n", qparams)

	qr, err := json.Marshal(qresp)
	if err != nil {
		return shim.Error("[getCOCampaigns] Error unable to Marshall qresp: " + err.Error())
	}

	fmt.Println("- end getCOCampaigns")
	return shim.Success(qr)
}
