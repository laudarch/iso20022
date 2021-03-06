package iso20022

// Point of interaction (POI) performing the transaction.
type PointOfInteraction2 struct {

	// Identification of the POI (Point Of Interaction) for the acquirer or its agent.
	Identification *GenericIdentification32 `xml:"Id"`

	// Common name assigned by the acquirer to the POI system.
	SystemName *Max70Text `xml:"SysNm,omitempty"`

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	GroupIdentification *Max35Text `xml:"GrpId,omitempty"`

	// Capabilities of the POI performing the transaction.
	Capabilities *PointOfInteractionCapabilities1 `xml:"Cpblties,omitempty"`

	// Data related to a component of the POI performing the transaction.
	Component []*PointOfInteractionComponent3 `xml:"Cmpnt,omitempty"`

}


func (p *PointOfInteraction2) AddIdentification() *GenericIdentification32 {
	p.Identification = new(GenericIdentification32)
	return p.Identification
}

func (p *PointOfInteraction2) SetSystemName(value string) {
	p.SystemName = (*Max70Text)(&value)
}

func (p *PointOfInteraction2) SetGroupIdentification(value string) {
	p.GroupIdentification = (*Max35Text)(&value)
}

func (p *PointOfInteraction2) AddCapabilities() *PointOfInteractionCapabilities1 {
	p.Capabilities = new(PointOfInteractionCapabilities1)
	return p.Capabilities
}

func (p *PointOfInteraction2) AddComponent() *PointOfInteractionComponent3 {
	newValue := new (PointOfInteractionComponent3)
	p.Component = append(p.Component, newValue)
	return newValue
}

