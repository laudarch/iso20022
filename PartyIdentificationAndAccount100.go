package iso20022

// Party and account details.
type PartyIdentificationAndAccount100 struct {

	// Identification of the party.
	Identification *PartyIdentification83Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party,
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`

}


func (p *PartyIdentificationAndAccount100) AddIdentification() *PartyIdentification83Choice {
	p.Identification = new(PartyIdentification83Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount100) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount100) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount100) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount100) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}

