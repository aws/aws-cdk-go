package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@AS2PayloadSendFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2PayloadSendFailed := #error#.NewAS2PayloadSendFailed()
//
// Experimental.
type ConnectorEvents_AS2PayloadSendFailed interface {
}

// The jsii proxy struct for ConnectorEvents_AS2PayloadSendFailed
type jsiiProxy_ConnectorEvents_AS2PayloadSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_AS2PayloadSendFailed() ConnectorEvents_AS2PayloadSendFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_AS2PayloadSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_AS2PayloadSendFailed_Override(c ConnectorEvents_AS2PayloadSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.AS2PayloadSendFailed",
		nil, // no parameters
		c,
	)
}

