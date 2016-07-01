package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00600107 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.006.001.07 Document"`
	Message *TransferInCancellationRequestV07 `xml:"TrfInCxlReq"`
}

func (d *Document00600107) AddMessage() *TransferInCancellationRequestV07 {
	d.Message = new(TransferInCancellationRequestV07)
	return d.Message
}

// Scope
// An instructing party, for example, a transfer agent, sends the TransferInCancellationRequest message to the executing party, for example, a transfer agent, to request the cancellation of a previously sent TransferInInstruction.
// Usage
// The TransferInCancellationRequest message is used to request cancellation of a previously sent TransferInInstruction.
// There are two ways to specify the transfer cancellation request. Either:
// - the transfer reference of the original transfer is quoted, or,
// - all the details of the original transfer (this includes TransferReference) are quoted but this is not recommended.
// The message identification of the TransferInInstruction message in which the transfer was conveyed may also be quoted in PreviousReference. It is also possible to request the cancellation of a TransferInInstruction message by quoting its message identification in PreviousReference.
type TransferInCancellationRequestV07 struct {

	// Reference that uniquely identifies a message from a business application standpoint.
	MessageIdentification *iso20022.MessageIdentification1 `xml:"MsgId"`

	// Reference to the transaction identifier issued by the counterparty. Building block may also be used to reference a previous transaction, or tie a set of messages together.
	References []*iso20022.References20 `xml:"Refs,omitempty"`

	// Function of the transfer-in, that is, whether the message is used as a request to cancel a previously sent instruction or as a cancellation of a previously sent advice and request for information. The absence of Function indicates the message is a request to cancel a previously sent instruction.
	Function *iso20022.TransferInFunction1Code `xml:"Fctn,omitempty"`

	// Choice between cancellation by reference or by transfer details.
	Cancellation []*iso20022.Cancellation10Choice `xml:"Cxl"`

	// Identifies the market practice to which the message conforms.
	MarketPracticeVersion *iso20022.MarketPracticeVersion1 `xml:"MktPrctcVrsn,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *iso20022.CopyInformation4 `xml:"CpyDtls,omitempty"`

}


func (t *TransferInCancellationRequestV07) AddMessageIdentification() *iso20022.MessageIdentification1 {
	t.MessageIdentification = new(iso20022.MessageIdentification1)
	return t.MessageIdentification
}

func (t *TransferInCancellationRequestV07) AddReferences() *iso20022.References20 {
	newValue := new (iso20022.References20)
	t.References = append(t.References, newValue)
	return newValue
}

func (t *TransferInCancellationRequestV07) SetFunction(value string) {
	t.Function = (*iso20022.TransferInFunction1Code)(&value)
}

func (t *TransferInCancellationRequestV07) AddCancellation() *iso20022.Cancellation10Choice {
	newValue := new (iso20022.Cancellation10Choice)
	t.Cancellation = append(t.Cancellation, newValue)
	return newValue
}

func (t *TransferInCancellationRequestV07) AddMarketPracticeVersion() *iso20022.MarketPracticeVersion1 {
	t.MarketPracticeVersion = new(iso20022.MarketPracticeVersion1)
	return t.MarketPracticeVersion
}

func (t *TransferInCancellationRequestV07) AddCopyDetails() *iso20022.CopyInformation4 {
	t.CopyDetails = new(iso20022.CopyInformation4)
	return t.CopyDetails
}

