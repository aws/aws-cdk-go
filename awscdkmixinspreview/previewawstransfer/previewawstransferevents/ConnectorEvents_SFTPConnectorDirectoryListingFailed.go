package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorDirectoryListingFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingFailed := #error#.NewSFTPConnectorDirectoryListingFailed()
//
// Experimental.
type ConnectorEvents_SFTPConnectorDirectoryListingFailed interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorDirectoryListingFailed
type jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorDirectoryListingFailed() ConnectorEvents_SFTPConnectorDirectoryListingFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorDirectoryListingFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorDirectoryListingFailed_Override(c ConnectorEvents_SFTPConnectorDirectoryListingFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorDirectoryListingFailed",
		nil, // no parameters
		c,
	)
}

