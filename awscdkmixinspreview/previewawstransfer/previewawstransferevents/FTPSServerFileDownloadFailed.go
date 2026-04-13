package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileDownloadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileDownloadFailed := awscdkmixinspreview.Events.NewFTPSServerFileDownloadFailed()
//
// Experimental.
type FTPSServerFileDownloadFailed interface {
}

// The jsii proxy struct for FTPSServerFileDownloadFailed
type jsiiProxy_FTPSServerFileDownloadFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileDownloadFailed() FTPSServerFileDownloadFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileDownloadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileDownloadFailed_Override(f FTPSServerFileDownloadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Download Failed.
// Experimental.
func FTPSServerFileDownloadFailed_EventPattern(options *FTPSServerFileDownloadFailed_FTPSServerFileDownloadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileDownloadFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

