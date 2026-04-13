package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileUploadCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileUploadCompleted := awscdkmixinspreview.Events.NewFTPServerFileUploadCompleted()
//
// Experimental.
type FTPServerFileUploadCompleted interface {
}

// The jsii proxy struct for FTPServerFileUploadCompleted
type jsiiProxy_FTPServerFileUploadCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileUploadCompleted() FTPServerFileUploadCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileUploadCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileUploadCompleted_Override(f FTPServerFileUploadCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Upload Completed.
// Experimental.
func FTPServerFileUploadCompleted_EventPattern(options *FTPServerFileUploadCompleted_FTPServerFileUploadCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileUploadCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileUploadCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

