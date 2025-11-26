package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorFileSendFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileSendFailed := #error#.NewSFTPConnectorFileSendFailed()
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileSendFailed interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorFileSendFailed
type jsiiProxy_ConnectorEvents_SFTPConnectorFileSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileSendFailed() ConnectorEvents_SFTPConnectorFileSendFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorFileSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileSendFailed_Override(c ConnectorEvents_SFTPConnectorFileSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileSendFailed",
		nil, // no parameters
		c,
	)
}

