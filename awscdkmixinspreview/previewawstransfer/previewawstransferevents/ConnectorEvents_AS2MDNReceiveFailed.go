package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@AS2MDNReceiveFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNReceiveFailed := #error#.NewAS2MDNReceiveFailed()
//
// Experimental.
type ConnectorEvents_AS2MDNReceiveFailed interface {
}

// The jsii proxy struct for ConnectorEvents_AS2MDNReceiveFailed
type jsiiProxy_ConnectorEvents_AS2MDNReceiveFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_AS2MDNReceiveFailed() ConnectorEvents_AS2MDNReceiveFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_AS2MDNReceiveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_AS2MDNReceiveFailed_Override(c ConnectorEvents_AS2MDNReceiveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveFailed",
		nil, // no parameters
		c,
	)
}

