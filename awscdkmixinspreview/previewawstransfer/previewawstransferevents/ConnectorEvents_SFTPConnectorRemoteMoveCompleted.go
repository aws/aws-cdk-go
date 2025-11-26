package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorRemoteMoveCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteMoveCompleted := #error#.NewSFTPConnectorRemoteMoveCompleted()
//
// Experimental.
type ConnectorEvents_SFTPConnectorRemoteMoveCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorRemoteMoveCompleted
type jsiiProxy_ConnectorEvents_SFTPConnectorRemoteMoveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteMoveCompleted() ConnectorEvents_SFTPConnectorRemoteMoveCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorRemoteMoveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorRemoteMoveCompleted_Override(c ConnectorEvents_SFTPConnectorRemoteMoveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorRemoteMoveCompleted",
		nil, // no parameters
		c,
	)
}

