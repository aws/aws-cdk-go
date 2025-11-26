package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.transfer@SFTPConnectorFileRetrieveCompleted event types for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileRetrieveCompleted := #error#.NewSFTPConnectorFileRetrieveCompleted()
//
// Experimental.
type ConnectorEvents_SFTPConnectorFileRetrieveCompleted interface {
}

// The jsii proxy struct for ConnectorEvents_SFTPConnectorFileRetrieveCompleted
type jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileRetrieveCompleted() ConnectorEvents_SFTPConnectorFileRetrieveCompleted {
	_init_.Initialize()

	j := jsiiProxy_ConnectorEvents_SFTPConnectorFileRetrieveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConnectorEvents_SFTPConnectorFileRetrieveCompleted_Override(c ConnectorEvents_SFTPConnectorFileRetrieveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents.SFTPConnectorFileRetrieveCompleted",
		nil, // no parameters
		c,
	)
}

