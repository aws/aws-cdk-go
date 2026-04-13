package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorFileSendCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileSendCompleted := awscdkmixinspreview.Events.NewSFTPConnectorFileSendCompleted()
//
// Experimental.
type SFTPConnectorFileSendCompleted interface {
}

// The jsii proxy struct for SFTPConnectorFileSendCompleted
type jsiiProxy_SFTPConnectorFileSendCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorFileSendCompleted() SFTPConnectorFileSendCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorFileSendCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorFileSendCompleted_Override(s SFTPConnectorFileSendCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector File Send Completed.
// Experimental.
func SFTPConnectorFileSendCompleted_EventPattern(options *SFTPConnectorFileSendCompleted_SFTPConnectorFileSendCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorFileSendCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

