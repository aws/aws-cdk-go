package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileDownloadCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileDownloadCompleted := awscdkmixinspreview.Events.NewFTPServerFileDownloadCompleted()
//
// Experimental.
type FTPServerFileDownloadCompleted interface {
}

// The jsii proxy struct for FTPServerFileDownloadCompleted
type jsiiProxy_FTPServerFileDownloadCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileDownloadCompleted() FTPServerFileDownloadCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileDownloadCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileDownloadCompleted_Override(f FTPServerFileDownloadCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Download Completed.
// Experimental.
func FTPServerFileDownloadCompleted_EventPattern(options *FTPServerFileDownloadCompleted_FTPServerFileDownloadCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileDownloadCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

