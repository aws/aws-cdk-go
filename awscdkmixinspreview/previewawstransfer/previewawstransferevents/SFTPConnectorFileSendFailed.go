package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorFileSendFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileSendFailed := awscdkmixinspreview.Events.NewSFTPConnectorFileSendFailed()
//
// Experimental.
type SFTPConnectorFileSendFailed interface {
}

// The jsii proxy struct for SFTPConnectorFileSendFailed
type jsiiProxy_SFTPConnectorFileSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorFileSendFailed() SFTPConnectorFileSendFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorFileSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorFileSendFailed_Override(s SFTPConnectorFileSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector File Send Failed.
// Experimental.
func SFTPConnectorFileSendFailed_SftPConnectorFileSendFailedPattern(options *SFTPConnectorFileSendFailed_SFTPConnectorFileSendFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorFileSendFailed_SftPConnectorFileSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileSendFailed",
		"sftPConnectorFileSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

