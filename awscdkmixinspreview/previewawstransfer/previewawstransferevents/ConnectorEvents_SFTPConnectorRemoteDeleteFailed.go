package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorRemoteDeleteFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteDeleteFailed := #error#.NewSFTPConnectorRemoteDeleteFailed()
//
// Experimental.
type ConnectorEvents_SFTPConnectorRemoteDeleteFailed interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorRemoteDeleteFailed
type jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteDeleteFailed() ConnectorEvents_SFTPConnectorRemoteDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteDeleteFailed_Override(c ConnectorEvents_SFTPConnectorRemoteDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteFailed",
		nil, // no parameters
		c,
	)
}

