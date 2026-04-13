package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileUploadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileUploadFailed := awscdkmixinspreview.Events.NewFTPSServerFileUploadFailed()
//
// Experimental.
type FTPSServerFileUploadFailed interface {
}

// The jsii proxy struct for FTPSServerFileUploadFailed
type jsiiProxy_FTPSServerFileUploadFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileUploadFailed() FTPSServerFileUploadFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileUploadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileUploadFailed_Override(f FTPSServerFileUploadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Upload Failed.
// Experimental.
func FTPSServerFileUploadFailed_EventPattern(options *FTPSServerFileUploadFailed_FTPSServerFileUploadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileUploadFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

