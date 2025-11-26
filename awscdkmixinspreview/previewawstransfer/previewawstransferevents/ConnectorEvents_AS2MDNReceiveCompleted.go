package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@AS2MDNReceiveCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNReceiveCompleted := #error#.NewAS2MDNReceiveCompleted()
//
// Experimental.
type ConnectorEvents_AS2MDNReceiveCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_AS2MDNReceiveCompleted
type jsiiProxy_ConnectorEvents_AS2MDNReceiveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_AS2MDNReceiveCompleted() ConnectorEvents_AS2MDNReceiveCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_AS2MDNReceiveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_AS2MDNReceiveCompleted_Override(c ConnectorEvents_AS2MDNReceiveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2MDNReceiveCompleted",
		nil, // no parameters
		c,
	)
}

