package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorRemoteDeleteCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteDeleteCompleted := #error#.NewSFTPConnectorRemoteDeleteCompleted()
//
// Experimental.
type ConnectorEvents_SFTPConnectorRemoteDeleteCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorRemoteDeleteCompleted
type jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteDeleteCompleted() ConnectorEvents_SFTPConnectorRemoteDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorRemoteDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteDeleteCompleted_Override(c ConnectorEvents_SFTPConnectorRemoteDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteDeleteCompleted",
		nil, // no parameters
		c,
	)
}

