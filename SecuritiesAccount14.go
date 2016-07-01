package iso20022

// Account to or from which a securities entry is made.
type SecuritiesAccount14 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode2Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`

}


func (s *SecuritiesAccount14) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount14) AddType() *PurposeCode2Choice {
	s.Type = new(PurposeCode2Choice)
	return s.Type
}

func (s *SecuritiesAccount14) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

