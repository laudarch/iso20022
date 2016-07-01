package iso20022

// Set of characteristics that apply to the debit side of the payment transactions included in the credit transfer transaction initiation.
type PaymentInstructionInformation1 struct {

	// Reference assigned by a sending party to unambiguously identify the payment information block within the message.
	PaymentInformationIdentification *Max35Text `xml:"PmtInfId,omitempty"`

	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod *PaymentMethod3Code `xml:"PmtMtd"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation1 `xml:"PmtTpInf,omitempty"`

	// Date at which the initiating party requests the clearing agent to process the payment. If payment by cheque, the date when the cheque must be generated by the bank.
	// 
	// Usage: This is the date on which the debtor's account(s) is (are) to be debited.
	RequestedExecutionDate *ISODate `xml:"ReqdExctnDt"`

	// Date used for the correction of the value date of a cash pool movement that has been posted with a different value date.
	PoolingAdjustmentDate *ISODate `xml:"PoolgAdjstmntDt,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification8 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount7 `xml:"DbtrAcct"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification3 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount7 `xml:"DbtrAgtAcct,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification8 `xml:"UltmtDbtr,omitempty"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Account used to process charges associated with a transaction.
	// 
	// Usage : charges account should be used when charges have to be booked to an account different from the account identified in debtor's account.
	ChargesAccount *CashAccount7 `xml:"ChrgsAcct,omitempty"`

	// Agent that services a charges account.  
	// 
	// Usage : charges account agent should only be used when the charges account agent is different from the debtor agent.
	ChargesAccountAgent *BranchAndFinancialInstitutionIdentification3 `xml:"ChrgsAcctAgt,omitempty"`

	// Set of elements providing information specific to the individual transaction(s) included in the message.
	CreditTransferTransactionInformation []*CreditTransferTransactionInformation1 `xml:"CdtTrfTxInf"`

}


func (p *PaymentInstructionInformation1) SetPaymentInformationIdentification(value string) {
	p.PaymentInformationIdentification = (*Max35Text)(&value)
}

func (p *PaymentInstructionInformation1) SetPaymentMethod(value string) {
	p.PaymentMethod = (*PaymentMethod3Code)(&value)
}

func (p *PaymentInstructionInformation1) AddPaymentTypeInformation() *PaymentTypeInformation1 {
	p.PaymentTypeInformation = new(PaymentTypeInformation1)
	return p.PaymentTypeInformation
}

func (p *PaymentInstructionInformation1) SetRequestedExecutionDate(value string) {
	p.RequestedExecutionDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation1) SetPoolingAdjustmentDate(value string) {
	p.PoolingAdjustmentDate = (*ISODate)(&value)
}

func (p *PaymentInstructionInformation1) AddDebtor() *PartyIdentification8 {
	p.Debtor = new(PartyIdentification8)
	return p.Debtor
}

func (p *PaymentInstructionInformation1) AddDebtorAccount() *CashAccount7 {
	p.DebtorAccount = new(CashAccount7)
	return p.DebtorAccount
}

func (p *PaymentInstructionInformation1) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.DebtorAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.DebtorAgent
}

func (p *PaymentInstructionInformation1) AddDebtorAgentAccount() *CashAccount7 {
	p.DebtorAgentAccount = new(CashAccount7)
	return p.DebtorAgentAccount
}

func (p *PaymentInstructionInformation1) AddUltimateDebtor() *PartyIdentification8 {
	p.UltimateDebtor = new(PartyIdentification8)
	return p.UltimateDebtor
}

func (p *PaymentInstructionInformation1) SetChargeBearer(value string) {
	p.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (p *PaymentInstructionInformation1) AddChargesAccount() *CashAccount7 {
	p.ChargesAccount = new(CashAccount7)
	return p.ChargesAccount
}

func (p *PaymentInstructionInformation1) AddChargesAccountAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.ChargesAccountAgent = new(BranchAndFinancialInstitutionIdentification3)
	return p.ChargesAccountAgent
}

func (p *PaymentInstructionInformation1) AddCreditTransferTransactionInformation() *CreditTransferTransactionInformation1 {
	newValue := new (CreditTransferTransactionInformation1)
	p.CreditTransferTransactionInformation = append(p.CreditTransferTransactionInformation, newValue)
	return newValue
}

