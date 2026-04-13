package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileDownloadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileDownloadFailed := awscdkmixinspreview.Events.NewFTPServerFileDownloadFailed()
//
// Experimental.
type FTPServerFileDownloadFailed interface {
}

// The jsii proxy struct for FTPServerFileDownloadFailed
type jsiiProxy_FTPServerFileDownloadFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileDownloadFailed() FTPServerFileDownloadFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileDownloadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileDownloadFailed_Override(f FTPServerFileDownloadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Download Failed.
// Experimental.
func FTPServerFileDownloadFailed_EventPattern(options *FTPServerFileDownloadFailed_FTPServerFileDownloadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileDownloadFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDownloadFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

