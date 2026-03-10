package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileUploadCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileUploadCompleted := awscdkmixinspreview.Events.NewFTPSServerFileUploadCompleted()
//
// Experimental.
type FTPSServerFileUploadCompleted interface {
}

// The jsii proxy struct for FTPSServerFileUploadCompleted
type jsiiProxy_FTPSServerFileUploadCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileUploadCompleted() FTPSServerFileUploadCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileUploadCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileUploadCompleted_Override(f FTPSServerFileUploadCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Upload Completed.
// Experimental.
func FTPSServerFileUploadCompleted_FtpSServerFileUploadCompletedPattern(options *FTPSServerFileUploadCompleted_FTPSServerFileUploadCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileUploadCompleted_FtpSServerFileUploadCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileUploadCompleted",
		"ftpSServerFileUploadCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

