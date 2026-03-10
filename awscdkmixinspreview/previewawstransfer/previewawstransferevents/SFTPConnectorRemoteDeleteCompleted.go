package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorRemoteDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteDeleteCompleted := awscdkmixinspreview.Events.NewSFTPConnectorRemoteDeleteCompleted()
//
// Experimental.
type SFTPConnectorRemoteDeleteCompleted interface {
}

// The jsii proxy struct for SFTPConnectorRemoteDeleteCompleted
type jsiiProxy_SFTPConnectorRemoteDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorRemoteDeleteCompleted() SFTPConnectorRemoteDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorRemoteDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorRemoteDeleteCompleted_Override(s SFTPConnectorRemoteDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Remote Delete Completed.
// Experimental.
func SFTPConnectorRemoteDeleteCompleted_SftPConnectorRemoteDeleteCompletedPattern(options *SFTPConnectorRemoteDeleteCompleted_SFTPConnectorRemoteDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorRemoteDeleteCompleted_SftPConnectorRemoteDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteCompleted",
		"sftPConnectorRemoteDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

