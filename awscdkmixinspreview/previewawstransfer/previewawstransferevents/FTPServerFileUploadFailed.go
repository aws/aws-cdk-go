package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileUploadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileUploadFailed := awscdkmixinspreview.Events.NewFTPServerFileUploadFailed()
//
// Experimental.
type FTPServerFileUploadFailed interface {
}

// The jsii proxy struct for FTPServerFileUploadFailed
type jsiiProxy_FTPServerFileUploadFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileUploadFailed() FTPServerFileUploadFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileUploadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileUploadFailed_Override(f FTPServerFileUploadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Upload Failed.
// Experimental.
func FTPServerFileUploadFailed_EventPattern(options *FTPServerFileUploadFailed_FTPServerFileUploadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileUploadFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

