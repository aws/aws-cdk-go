package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerDirectoryCreateFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerDirectoryCreateFailed := awscdkmixinspreview.Events.NewSFTPServerDirectoryCreateFailed()
//
// Experimental.
type SFTPServerDirectoryCreateFailed interface {
}

// The jsii proxy struct for SFTPServerDirectoryCreateFailed
type jsiiProxy_SFTPServerDirectoryCreateFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerDirectoryCreateFailed() SFTPServerDirectoryCreateFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerDirectoryCreateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerDirectoryCreateFailed_Override(s SFTPServerDirectoryCreateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server Directory Create Failed.
// Experimental.
func SFTPServerDirectoryCreateFailed_EventPattern(options *SFTPServerDirectoryCreateFailed_SFTPServerDirectoryCreateFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerDirectoryCreateFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

