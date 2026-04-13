package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileDeleteFailed := awscdkmixinspreview.Events.NewSFTPServerFileDeleteFailed()
//
// Experimental.
type SFTPServerFileDeleteFailed interface {
}

// The jsii proxy struct for SFTPServerFileDeleteFailed
type jsiiProxy_SFTPServerFileDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileDeleteFailed() SFTPServerFileDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileDeleteFailed_Override(s SFTPServerFileDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Delete Failed.
// Experimental.
func SFTPServerFileDeleteFailed_EventPattern(options *SFTPServerFileDeleteFailed_SFTPServerFileDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileDeleteFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

