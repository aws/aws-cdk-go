package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileDeleteCompleted := awscdkmixinspreview.Events.NewFTPServerFileDeleteCompleted()
//
// Experimental.
type FTPServerFileDeleteCompleted interface {
}

// The jsii proxy struct for FTPServerFileDeleteCompleted
type jsiiProxy_FTPServerFileDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileDeleteCompleted() FTPServerFileDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileDeleteCompleted_Override(f FTPServerFileDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Delete Completed.
// Experimental.
func FTPServerFileDeleteCompleted_FtpServerFileDeleteCompletedPattern(options *FTPServerFileDeleteCompleted_FTPServerFileDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileDeleteCompleted_FtpServerFileDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteCompleted",
		"ftpServerFileDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

