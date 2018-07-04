package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func createQueryString(params *CampaignParams) (qstring string, err error) {
	//ex: {"selector":{"CharityID":"marble","Status":1}
	var selector = CampaignQuerySelector{Selector: *params}
	serialized, err := json.Marshal(selector)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	qstring = string(serialized)
	return qstring, nil
}

func applyIdentityModsFromParam(src *CampaignParams, dest *CampaignInfo) {
	if "" != src.DonatedAmount {
		dstamt, _ := strconv.ParseFloat(dest.DonatedAmount, 64)
		srcamt, _ := strconv.ParseFloat(src.DonatedAmount, 64)
		dest.DonatedAmount = FloatToString(srcamt + dstamt)
	}
	if "" != src.Status {
		dest.Status = src.Status
	}
	if "" != src.DisbursedAmount {
		dstamt, _ := strconv.ParseFloat(dest.DisbursedAmount, 64)
		srcamt, _ := strconv.ParseFloat(src.DisbursedAmount, 64)
		dest.DisbursedAmount = FloatToString(srcamt + dstamt)
	}
	if 0 != src.RatingFive {
		dest.RatingFive++
	}
	if 0 != src.RatingFour {
		dest.RatingFour++
	}
	if 0 != src.RatingThree {
		dest.RatingThree++
	}
	if 0 != src.RatingTwo {
		dest.RatingTwo++
	}
	if 0 != src.RatingOne {
		dest.RatingOne++
	}
}
