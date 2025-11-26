package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorDirectoryListingCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingCompleted := #error#.NewSFTPConnectorDirectoryListingCompleted()
//
// Experimental.
type ConnectorEvents_SFTPConnectorDirectoryListingCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorDirectoryListingCompleted
type jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorDirectoryListingCompleted() ConnectorEvents_SFTPConnectorDirectoryListingCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorDirectoryListingCompleted_Override(c ConnectorEvents_SFTPConnectorDirectoryListingCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingCompleted",
		nil, // no parameters
		c,
	)
}

