package iso20022

// Choice of format for the rejection reason.
type RejectionReason15Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason38Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`

}


func (r *RejectionReason15Choice) SetCode(value string) {
	r.Code = (*RejectionReason38Code)(&value)
}

func (r *RejectionReason15Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}

