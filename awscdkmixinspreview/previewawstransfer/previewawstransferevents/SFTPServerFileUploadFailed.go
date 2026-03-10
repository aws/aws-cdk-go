package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileUploadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileUploadFailed := awscdkmixinspreview.Events.NewSFTPServerFileUploadFailed()
//
// Experimental.
type SFTPServerFileUploadFailed interface {
}

// The jsii proxy struct for SFTPServerFileUploadFailed
type jsiiProxy_SFTPServerFileUploadFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileUploadFailed() SFTPServerFileUploadFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileUploadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileUploadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileUploadFailed_Override(s SFTPServerFileUploadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileUploadFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Upload Failed.
// Experimental.
func SFTPServerFileUploadFailed_SftPServerFileUploadFailedPattern(options *SFTPServerFileUploadFailed_SFTPServerFileUploadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileUploadFailed_SftPServerFileUploadFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileUploadFailed",
		"sftPServerFileUploadFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

