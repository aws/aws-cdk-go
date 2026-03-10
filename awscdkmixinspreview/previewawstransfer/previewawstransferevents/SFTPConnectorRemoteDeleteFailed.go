package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorRemoteDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteDeleteFailed := awscdkmixinspreview.Events.NewSFTPConnectorRemoteDeleteFailed()
//
// Experimental.
type SFTPConnectorRemoteDeleteFailed interface {
}

// The jsii proxy struct for SFTPConnectorRemoteDeleteFailed
type jsiiProxy_SFTPConnectorRemoteDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorRemoteDeleteFailed() SFTPConnectorRemoteDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorRemoteDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorRemoteDeleteFailed_Override(s SFTPConnectorRemoteDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Remote Delete Failed.
// Experimental.
func SFTPConnectorRemoteDeleteFailed_SftPConnectorRemoteDeleteFailedPattern(options *SFTPConnectorRemoteDeleteFailed_SFTPConnectorRemoteDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorRemoteDeleteFailed_SftPConnectorRemoteDeleteFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteDeleteFailed",
		"sftPConnectorRemoteDeleteFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

