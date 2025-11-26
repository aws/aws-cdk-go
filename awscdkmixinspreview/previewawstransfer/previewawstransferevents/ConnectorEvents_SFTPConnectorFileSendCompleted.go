package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorFileSendCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileSendCompleted := #error#.NewSFTPConnectorFileSendCompleted()
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileSendCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorFileSendCompleted
type jsiiProxy_ConnectorEvents_SFTPConnectorFileSendCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileSendCompleted() ConnectorEvents_SFTPConnectorFileSendCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorFileSendCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileSendCompleted_Override(c ConnectorEvents_SFTPConnectorFileSendCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendCompleted",
		nil, // no parameters
		c,
	)
}

