package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorFileRetrieveFailed event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileRetrieveFailed := #error#.NewSFTPConnectorFileRetrieveFailed()
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileRetrieveFailed interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorFileRetrieveFailed
type jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveFailed struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileRetrieveFailed() ConnectorEvents_SFTPConnectorFileRetrieveFailed {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileRetrieveFailed_Override(c ConnectorEvents_SFTPConnectorFileRetrieveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveFailed",
		nil, // no parameters
		c,
	)
}

