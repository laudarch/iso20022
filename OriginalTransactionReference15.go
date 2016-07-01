package iso20022

// Set of key elements used to refer the original transaction.
type OriginalTransactionReference15 struct {

	// Amount of money moved between the instructing agent and the instructed agent.
	InterbankSettlementAmount *ActiveOrHistoricCurrencyAndAmount `xml:"IntrBkSttlmAmt,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	Amount *AmountType3Choice `xml:"Amt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Date and time at which the creditor requests that the amount of money is to be collected from the debtor.
	RequestedCollectionDate *ISODate `xml:"ReqdColltnDt,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. 
	// Usage: This is the date on which the debtor's account is to be debited. If payment by cheque, the date when the cheque must be generated by the bank.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Specifies the details on how the settlement of the original transaction(s) between the instructing agent and the instructed agent was completed.
	SettlementInformation *SettlementInformation16 `xml:"SttlmInf,omitempty"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation22 `xml:"PmtTpInf,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod4Code `xml:"PmtMtd,omitempty"`

	// Set of elements used to provide further details of the mandate signed between the creditor and the debtor.
	MandateRelatedInformation *MandateRelatedInformation7 `xml:"MndtRltdInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation6 `xml:"RmtInf,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount16 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt"`

	// Party to which an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount16 `xml:"CdtrAcct,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

}


func (o *OriginalTransactionReference15) SetInterbankSettlementAmount(value, currency string) {
	o.InterbankSettlementAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (o *OriginalTransactionReference15) AddAmount() *AmountType3Choice {
	o.Amount = new(AmountType3Choice)
	return o.Amount
}

func (o *OriginalTransactionReference15) SetInterbankSettlementDate(value string) {
	o.InterbankSettlementDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference15) SetRequestedCollectionDate(value string) {
	o.RequestedCollectionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference15) SetRequestedExecutionDate(value string) {
	o.RequestedExecutionDate = (*ISODate)(&value)
}

func (o *OriginalTransactionReference15) AddCreditorSchemeIdentification() *PartyIdentification43 {
	o.CreditorSchemeIdentification = new(PartyIdentification43)
	return o.CreditorSchemeIdentification
}

func (o *OriginalTransactionReference15) AddSettlementInformation() *SettlementInformation16 {
	o.SettlementInformation = new(SettlementInformation16)
	return o.SettlementInformation
}

func (o *OriginalTransactionReference15) AddPaymentTypeInformation() *PaymentTypeInformation22 {
	o.PaymentTypeInformation = new(PaymentTypeInformation22)
	return o.PaymentTypeInformation
}

func (o *OriginalTransactionReference15) SetPaymentMethod(value string) {
	o.PaymentMethod = (*PaymentMethod4Code)(&value)
}

func (o *OriginalTransactionReference15) AddMandateRelatedInformation() *MandateRelatedInformation7 {
	o.MandateRelatedInformation = new(MandateRelatedInformation7)
	return o.MandateRelatedInformation
}

func (o *OriginalTransactionReference15) AddRemittanceInformation() *RemittanceInformation6 {
	o.RemittanceInformation = new(RemittanceInformation6)
	return o.RemittanceInformation
}

func (o *OriginalTransactionReference15) AddUltimateDebtor() *PartyIdentification43 {
	o.UltimateDebtor = new(PartyIdentification43)
	return o.UltimateDebtor
}

func (o *OriginalTransactionReference15) AddDebtor() *PartyIdentification43 {
	o.Debtor = new(PartyIdentification43)
	return o.Debtor
}

func (o *OriginalTransactionReference15) AddDebtorAccount() *CashAccount16 {
	o.DebtorAccount = new(CashAccount16)
	return o.DebtorAccount
}

func (o *OriginalTransactionReference15) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.DebtorAgent
}

func (o *OriginalTransactionReference15) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	o.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return o.CreditorAgent
}

func (o *OriginalTransactionReference15) AddCreditor() *PartyIdentification43 {
	o.Creditor = new(PartyIdentification43)
	return o.Creditor
}

func (o *OriginalTransactionReference15) AddCreditorAccount() *CashAccount16 {
	o.CreditorAccount = new(CashAccount16)
	return o.CreditorAccount
}

func (o *OriginalTransactionReference15) AddUltimateCreditor() *PartyIdentification43 {
	o.UltimateCreditor = new(PartyIdentification43)
	return o.UltimateCreditor
}

