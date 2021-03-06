package iso20022

// Choice between a rate or an unspecified rate.
type RateFormat8Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Cash amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`

}


func (r *RateFormat8Choice) SetRate(value string) {
	r.Rate = (*PercentageRate)(&value)
}

func (r *RateFormat8Choice) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (r *RateFormat8Choice) SetNotSpecifiedRate(value string) {
	r.NotSpecifiedRate = (*RateValueType7Code)(&value)
}

