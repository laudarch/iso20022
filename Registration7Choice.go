package iso20022

// Choice of format for the registration information.
type Registration7Choice struct {

	// Registration information expressed as an ISO 20022 code.
	Code *Registration2Code `xml:"Cd"`

	// Registration information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`

}


func (r *Registration7Choice) SetCode(value string) {
	r.Code = (*Registration2Code)(&value)
}

func (r *Registration7Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}

