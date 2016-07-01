package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03200102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.032.001.02 Document"`
	Message *SecuritiesSettlementTransactionGenerationNotificationV02 `xml:"SctiesSttlmTxGnrtnNtfctn"`
}

func (d *Document03200102) AddMessage() *SecuritiesSettlementTransactionGenerationNotificationV02 {
	d.Message = new(SecuritiesSettlementTransactionGenerationNotificationV02)
	return d.Message
}

// Scope
// An account servicer sends a SecuritiesSettlementTransactionGenerationNotification to an account owner to advise the account owner of a securities settlement transaction that has been generated/created by the account servicer for the account owner. The reason for creation can range from the automatic transformation of pending settlement instructions following a corporate event to the recovery of an account owner transactions' database initiated by its account servicer.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// using the relevant elements in the Business Application Header.
// Usage
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information.
// using the relevant elements in the Business Application Header.
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type SecuritiesSettlementTransactionGenerationNotificationV02 struct {

	// Provides transaction type and identification information.
	TransactionIdentificationDetails *iso20022.SettlementTypeAndIdentification10 `xml:"TxIdDtls"`

	// Count of the number of transactions linked.
	NumberCounts *iso20022.NumberCount1Choice `xml:"NbCounts,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*iso20022.Linkages7 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *iso20022.SecuritiesTradeDetails1 `xml:"TradDtls"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *iso20022.SecurityIdentification14 `xml:"FinInstrmId"`

	// Elements characterising a financial instrument.
	FinancialInstrumentAttributes *iso20022.FinancialInstrumentAttributes20 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails []*iso20022.QuantityAndAccount17 `xml:"QtyAndAcctDtls"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *iso20022.SettlementDetails22 `xml:"SttlmParams"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *iso20022.SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *iso20022.SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *iso20022.CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *iso20022.AmountAndDirection2 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *iso20022.OtherAmounts12 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *iso20022.OtherParties8 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *iso20022.RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason []*iso20022.GeneratedReason1 `xml:"GnrtdRsn,omitempty"`

	// Status and reason of the transaction.
	StatusAndReason *iso20022.StatusAndReason3 `xml:"StsAndRsn,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddTransactionIdentificationDetails() *iso20022.SettlementTypeAndIdentification10 {
	s.TransactionIdentificationDetails = new(iso20022.SettlementTypeAndIdentification10)
	return s.TransactionIdentificationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddNumberCounts() *iso20022.NumberCount1Choice {
	s.NumberCounts = new(iso20022.NumberCount1Choice)
	return s.NumberCounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddLinkages() *iso20022.Linkages7 {
	newValue := new (iso20022.Linkages7)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddTradeDetails() *iso20022.SecuritiesTradeDetails1 {
	s.TradeDetails = new(iso20022.SecuritiesTradeDetails1)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddFinancialInstrumentIdentification() *iso20022.SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(iso20022.SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddFinancialInstrumentAttributes() *iso20022.FinancialInstrumentAttributes20 {
	s.FinancialInstrumentAttributes = new(iso20022.FinancialInstrumentAttributes20)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddQuantityAndAccountDetails() *iso20022.QuantityAndAccount17 {
	newValue := new (iso20022.QuantityAndAccount17)
	s.QuantityAndAccountDetails = append(s.QuantityAndAccountDetails, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddSettlementParameters() *iso20022.SettlementDetails22 {
	s.SettlementParameters = new(iso20022.SettlementDetails22)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddDeliveringSettlementParties() *iso20022.SettlementParties11 {
	s.DeliveringSettlementParties = new(iso20022.SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddReceivingSettlementParties() *iso20022.SettlementParties11 {
	s.ReceivingSettlementParties = new(iso20022.SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddCashParties() *iso20022.CashParties8 {
	s.CashParties = new(iso20022.CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddSettlementAmount() *iso20022.AmountAndDirection2 {
	s.SettlementAmount = new(iso20022.AmountAndDirection2)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddOtherAmounts() *iso20022.OtherAmounts12 {
	s.OtherAmounts = new(iso20022.OtherAmounts12)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddOtherBusinessParties() *iso20022.OtherParties8 {
	s.OtherBusinessParties = new(iso20022.OtherParties8)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddAdditionalPhysicalOrRegistrationDetails() *iso20022.RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(iso20022.RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddGeneratedReason() *iso20022.GeneratedReason1 {
	newValue := new (iso20022.GeneratedReason1)
	s.GeneratedReason = append(s.GeneratedReason, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddStatusAndReason() *iso20022.StatusAndReason3 {
	s.StatusAndReason = new(iso20022.StatusAndReason3)
	return s.StatusAndReason
}

func (s *SecuritiesSettlementTransactionGenerationNotificationV02) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}

