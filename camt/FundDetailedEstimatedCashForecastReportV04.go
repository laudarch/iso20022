package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04200104 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.042.001.04 Document"`
	Message *FundDetailedEstimatedCashForecastReportV04 `xml:"FndDtldEstmtdCshFcstRpt"`
}

func (d *Document04200104) AddMessage() *FundDetailedEstimatedCashForecastReportV04 {
	d.Message = new(FundDetailedEstimatedCashForecastReportV04)
	return d.Message
}

// Scope
// A report provider, such as a transfer agent, sends the FundDetailedEstimatedCashForecastReport message to the report user, such as an investment manager or pricing agent, to report the estimated cash incomings and outgoings, sorted by country, institution name or other criteria defined by the user of one or more share classes of an investment fund on one or more trade dates.
// The cash movements may result from, for example, redemption, subscription, switch transactions or reinvestment of dividends.
// Usage
// The FundDetailedEstimatedCashForecastReport is used to provide estimated cash movements, that is, it is sent prior to the cut-off time and/or the price valuation of the fund. The message contains incoming and outgoing cash flows that are estimated, that is, the price has not been applied. If the price is definitive, then the FundDetailedConfirmedCashForecastReport message must be used.
// The message structure allows for the following uses:
// -	to provide cash in and cash out amounts for a fund/sub fund and one or more share classes (a FundOrSubFundDetails sequence and one or more EstimatedFundCashForecastDetails sequences are used),
// -	to provide cash in and cash out amounts for one or more share classes (one or more EstimatedFundCashForecastDetails sequences are used).
// If the report is to provide estimated cash in and cash out for a fund/sub fund only and not for one or more share classes, then the FundEstimatedCashForecastReport message must be used.
// The FundDetailedEstimatedCashForecastReport message is used to report cash movements in or out of a fund, organised by party, such as fund management company, country, currency or by some other criteria defined by the report provider. If the report is used to give the cash-in and cash-out for a party, then additional criteria, such as currency and country, can be specified.
// In addition, the underlying transaction type for the cash-in or cash-out movement can be specified, as well as information about the cash movement's underlying orders, such as commission and charges.
type FundDetailedEstimatedCashForecastReportV04 struct {

	// Identifies the message.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Collective reference identifying a set of messages.
	PoolReference *iso20022.AdditionalReference3 `xml:"PoolRef,omitempty"`

	// Reference to a linked message that was previously sent.
	PreviousReference []*iso20022.AdditionalReference3 `xml:"PrvsRef,omitempty"`

	// Reference to a linked message that was previously received.
	RelatedReference []*iso20022.AdditionalReference3 `xml:"RltdRef,omitempty"`

	// Pagination of the message.
	MessagePagination *iso20022.Pagination `xml:"MsgPgntn"`

	// Information about the fund/sub fund when the report either specifies cash flow for the fund/sub fund or for a share class of the fund/sub fund.
	FundOrSubFundDetails *iso20022.Fund3 `xml:"FndOrSubFndDtls,omitempty"`

	// Information related to the estimated cash-in and cash-out flows for a specific trade date as a result of transactions in shares in an investment fund, for example, subscriptions, redemptions or switches. The information provided is sorted by pre-defined criteria such as country, institution, currency or user defined criteria.
	EstimatedFundCashForecastDetails []*iso20022.EstimatedFundCashForecast5 `xml:"EstmtdFndCshFcstDtls"`

	// Estimated net cash as a result of the cash-in and cash-out flows specified in the fund cash forecast details.
	ConsolidatedNetCashForecast *iso20022.NetCashForecast3 `xml:"CnsltdNetCshFcst,omitempty"`

	// Additional information that can not be captured in the structured fields and/or any other specific block. 
	Extension []*iso20022.Extension1 `xml:"Xtnsn,omitempty"`

}


func (f *FundDetailedEstimatedCashForecastReportV04) AddMessageIdentification() *iso20022.MessageIdentification1 {
	f.MessageIdentification = new(iso20022.MessageIdentification1)
	return f.MessageIdentification
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddPoolReference() *iso20022.AdditionalReference3 {
	f.PoolReference = new(iso20022.AdditionalReference3)
	return f.PoolReference
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddPreviousReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	f.PreviousReference = append(f.PreviousReference, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddRelatedReference() *iso20022.AdditionalReference3 {
	newValue := new (iso20022.AdditionalReference3)
	f.RelatedReference = append(f.RelatedReference, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddMessagePagination() *iso20022.Pagination {
	f.MessagePagination = new(iso20022.Pagination)
	return f.MessagePagination
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddFundOrSubFundDetails() *iso20022.Fund3 {
	f.FundOrSubFundDetails = new(iso20022.Fund3)
	return f.FundOrSubFundDetails
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddEstimatedFundCashForecastDetails() *iso20022.EstimatedFundCashForecast5 {
	newValue := new (iso20022.EstimatedFundCashForecast5)
	f.EstimatedFundCashForecastDetails = append(f.EstimatedFundCashForecastDetails, newValue)
	return newValue
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddConsolidatedNetCashForecast() *iso20022.NetCashForecast3 {
	f.ConsolidatedNetCashForecast = new(iso20022.NetCashForecast3)
	return f.ConsolidatedNetCashForecast
}

func (f *FundDetailedEstimatedCashForecastReportV04) AddExtension() *iso20022.Extension1 {
	newValue := new (iso20022.Extension1)
	f.Extension = append(f.Extension, newValue)
	return newValue
}

