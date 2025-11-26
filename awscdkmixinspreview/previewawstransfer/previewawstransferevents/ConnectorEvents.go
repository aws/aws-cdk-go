package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer"
)

// EventBridge event patterns for Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var connectorRef IConnectorRef
//
//   connectorEvents := awscdkmixinspreview.Events.ConnectorEvents_FromConnector(connectorRef)
//
// Experimental.
type ConnectorEvents interface {
	// EventBridge event pattern for Connector AS2 MDN Receive Completed.
	// Experimental.
	AS2MDNReceiveCompletedPattern(options *ConnectorEvents_AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 MDN Receive Failed.
	// Experimental.
	AS2MDNReceiveFailedPattern(options *ConnectorEvents_AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 Payload Send Completed.
	// Experimental.
	AS2PayloadSendCompletedPattern(options *ConnectorEvents_AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 Payload Send Failed.
	// Experimental.
	AS2PayloadSendFailedPattern(options *ConnectorEvents_AS2PayloadSendFailed_AS2PayloadSendFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Directory Listing Completed.
	// Experimental.
	SFTPConnectorDirectoryListingCompletedPattern(options *ConnectorEvents_SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Directory Listing Failed.
	// Experimental.
	SFTPConnectorDirectoryListingFailedPattern(options *ConnectorEvents_SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Retrieve Completed.
	// Experimental.
	SFTPConnectorFileRetrieveCompletedPattern(options *ConnectorEvents_SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Retrieve Failed.
	// Experimental.
	SFTPConnectorFileRetrieveFailedPattern(options *ConnectorEvents_SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Send Completed.
	// Experimental.
	SFTPConnectorFileSendCompletedPattern(options *ConnectorEvents_SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Send Failed.
	// Experimental.
	SFTPConnectorFileSendFailedPattern(options *ConnectorEvents_SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Delete Completed.
	// Experimental.
	SFTPConnectorRemoteDeleteCompletedPattern(options *ConnectorEvents_SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Delete Failed.
	// Experimental.
	SFTPConnectorRemoteDeleteFailedPattern(options *ConnectorEvents_SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Move Completed.
	// Experimental.
	SFTPConnectorRemoteMoveCompletedPattern(options *ConnectorEvents_SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Move Failed.
	// Experimental.
	SFTPConnectorRemoteMoveFailedPattern(options *ConnectorEvents_SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) *awsevents.EventPattern
}

// The jsii proxy struct for ConnectorEvents
type jsiiProxy_ConnectorEvents struct {
	_ byte // padding
}

// Create ConnectorEvents from a Connector reference.
// Experimental.
func ConnectorEvents_FromConnector(connectorRef interfacesawstransfer.IConnectorRef) ConnectorEvents {
	_init_.Initialize()

	if err := validateConnectorEvents_FromConnectorParameters(connectorRef); err != nil {
		panic(err)
	}
	var returns ConnectorEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.ConnectorEvents",
		"fromConnector",
		[]interface{}{connectorRef},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) AS2MDNReceiveCompletedPattern(options *ConnectorEvents_AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) *awsevents.EventPattern {
	if err := c.validateAS2MDNReceiveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"aS2MDNReceiveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) AS2MDNReceiveFailedPattern(options *ConnectorEvents_AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) *awsevents.EventPattern {
	if err := c.validateAS2MDNReceiveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"aS2MDNReceiveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) AS2PayloadSendCompletedPattern(options *ConnectorEvents_AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) *awsevents.EventPattern {
	if err := c.validateAS2PayloadSendCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"aS2PayloadSendCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) AS2PayloadSendFailedPattern(options *ConnectorEvents_AS2PayloadSendFailed_AS2PayloadSendFailedProps) *awsevents.EventPattern {
	if err := c.validateAS2PayloadSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"aS2PayloadSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorDirectoryListingCompletedPattern(options *ConnectorEvents_SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorDirectoryListingCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorDirectoryListingCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorDirectoryListingFailedPattern(options *ConnectorEvents_SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorDirectoryListingFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorDirectoryListingFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorFileRetrieveCompletedPattern(options *ConnectorEvents_SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorFileRetrieveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorFileRetrieveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorFileRetrieveFailedPattern(options *ConnectorEvents_SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorFileRetrieveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorFileRetrieveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorFileSendCompletedPattern(options *ConnectorEvents_SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorFileSendCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorFileSendCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorFileSendFailedPattern(options *ConnectorEvents_SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorFileSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorFileSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorRemoteDeleteCompletedPattern(options *ConnectorEvents_SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorRemoteDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorRemoteDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorRemoteDeleteFailedPattern(options *ConnectorEvents_SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorRemoteDeleteFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorRemoteDeleteFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorRemoteMoveCompletedPattern(options *ConnectorEvents_SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorRemoteMoveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorRemoteMoveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SFTPConnectorRemoteMoveFailedPattern(options *ConnectorEvents_SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) *awsevents.EventPattern {
	if err := c.validateSFTPConnectorRemoteMoveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sFTPConnectorRemoteMoveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

