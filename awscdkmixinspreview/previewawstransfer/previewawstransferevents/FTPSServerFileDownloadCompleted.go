package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileDownloadCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileDownloadCompleted := awscdkmixinspreview.Events.NewFTPSServerFileDownloadCompleted()
//
// Experimental.
type FTPSServerFileDownloadCompleted interface {
}

// The jsii proxy struct for FTPSServerFileDownloadCompleted
type jsiiProxy_FTPSServerFileDownloadCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileDownloadCompleted() FTPSServerFileDownloadCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileDownloadCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileDownloadCompleted_Override(f FTPSServerFileDownloadCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Download Completed.
// Experimental.
func FTPSServerFileDownloadCompleted_FtpSServerFileDownloadCompletedPattern(options *FTPSServerFileDownloadCompleted_FTPSServerFileDownloadCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileDownloadCompleted_FtpSServerFileDownloadCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDownloadCompleted",
		"ftpSServerFileDownloadCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

