package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@AS2MDNSendFailed event types for Agreement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNSendFailed := #error#.NewAS2MDNSendFailed()
//
// Experimental.
type AgreementEvents_AS2MDNSendFailed interface {
}

// The jsii proxy struct for AgreementEvents_AS2MDNSendFailed
type jsiiProxy_AgreementEvents_AS2MDNSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewAgreementEvents_AS2MDNSendFailed() AgreementEvents_AS2MDNSendFailed {
	_init_.Initialize()

	j := jsiiProxy_AgreementEvents_AS2MDNSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAgreementEvents_AS2MDNSendFailed_Override(a AgreementEvents_AS2MDNSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AgreementEvents.AS2MDNSendFailed",
		nil, // no parameters
		a,
	)
}

