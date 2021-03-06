package iso20022

// Amount of money debited or credited on the books of an account servicer.
type AmountAndDirection66 struct {

	// Amount of money in the cash entry.
	Amount *RestrictedFINActiveCurrencyAndAmount `xml:"Amt"`

	// Indicates whether an entry is a credit or a debit.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Posting/settlement amount in its original currency when conversion from/into another currency has occurred.
	OriginalCurrencyAndOrderedAmount *RestrictedFINActiveOrHistoricCurrencyAndAmount `xml:"OrgnlCcyAndOrdrdAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms27 `xml:"FXDtls,omitempty"`

}


func (a *AmountAndDirection66) SetAmount(value, currency string) {
	a.Amount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection66) SetCreditDebitIndicator(value string) {
	a.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (a *AmountAndDirection66) SetOriginalCurrencyAndOrderedAmount(value, currency string) {
	a.OriginalCurrencyAndOrderedAmount = NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection66) AddForeignExchangeDetails() *ForeignExchangeTerms27 {
	a.ForeignExchangeDetails = new(ForeignExchangeTerms27)
	return a.ForeignExchangeDetails
}

