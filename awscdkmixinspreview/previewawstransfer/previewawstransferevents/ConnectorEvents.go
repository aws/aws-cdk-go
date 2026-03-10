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
	As2MDNReceiveCompletedPattern(options *AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 MDN Receive Failed.
	// Experimental.
	As2MDNReceiveFailedPattern(options *AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 Payload Send Completed.
	// Experimental.
	As2PayloadSendCompletedPattern(options *AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector AS2 Payload Send Failed.
	// Experimental.
	As2PayloadSendFailedPattern(options *AS2PayloadSendFailed_AS2PayloadSendFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Directory Listing Completed.
	// Experimental.
	SftPConnectorDirectoryListingCompletedPattern(options *SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Directory Listing Failed.
	// Experimental.
	SftPConnectorDirectoryListingFailedPattern(options *SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Retrieve Completed.
	// Experimental.
	SftPConnectorFileRetrieveCompletedPattern(options *SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Retrieve Failed.
	// Experimental.
	SftPConnectorFileRetrieveFailedPattern(options *SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Send Completed.
	// Experimental.
	SftPConnectorFileSendCompletedPattern(options *SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector File Send Failed.
	// Experimental.
	SftPConnectorFileSendFailedPattern(options *SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Delete Completed.
	// Experimental.
	SftPConnectorRemoteDeleteCompletedPattern(options *SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Delete Failed.
	// Experimental.
	SftPConnectorRemoteDeleteFailedPattern(options *SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Move Completed.
	// Experimental.
	SftPConnectorRemoteMoveCompletedPattern(options *SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Connector SFTP Connector Remote Move Failed.
	// Experimental.
	SftPConnectorRemoteMoveFailedPattern(options *SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) *awsevents.EventPattern
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

func (c *jsiiProxy_ConnectorEvents) As2MDNReceiveCompletedPattern(options *AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) *awsevents.EventPattern {
	if err := c.validateAs2MDNReceiveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"as2MDNReceiveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) As2MDNReceiveFailedPattern(options *AS2MDNReceiveFailed_AS2MDNReceiveFailedProps) *awsevents.EventPattern {
	if err := c.validateAs2MDNReceiveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"as2MDNReceiveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) As2PayloadSendCompletedPattern(options *AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) *awsevents.EventPattern {
	if err := c.validateAs2PayloadSendCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"as2PayloadSendCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) As2PayloadSendFailedPattern(options *AS2PayloadSendFailed_AS2PayloadSendFailedProps) *awsevents.EventPattern {
	if err := c.validateAs2PayloadSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"as2PayloadSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorDirectoryListingCompletedPattern(options *SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorDirectoryListingCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorDirectoryListingCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorDirectoryListingFailedPattern(options *SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorDirectoryListingFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorDirectoryListingFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorFileRetrieveCompletedPattern(options *SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorFileRetrieveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorFileRetrieveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorFileRetrieveFailedPattern(options *SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorFileRetrieveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorFileRetrieveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorFileSendCompletedPattern(options *SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorFileSendCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorFileSendCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorFileSendFailedPattern(options *SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorFileSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorFileSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorRemoteDeleteCompletedPattern(options *SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorRemoteDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorRemoteDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorRemoteDeleteFailedPattern(options *SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorRemoteDeleteFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorRemoteDeleteFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorRemoteMoveCompletedPattern(options *SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorRemoteMoveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorRemoteMoveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectorEvents) SftPConnectorRemoteMoveFailedPattern(options *SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) *awsevents.EventPattern {
	if err := c.validateSftPConnectorRemoteMoveFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		c,
		"sftPConnectorRemoteMoveFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

