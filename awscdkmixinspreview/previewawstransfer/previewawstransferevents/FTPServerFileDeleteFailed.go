package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileDeleteFailed := awscdkmixinspreview.Events.NewFTPServerFileDeleteFailed()
//
// Experimental.
type FTPServerFileDeleteFailed interface {
}

// The jsii proxy struct for FTPServerFileDeleteFailed
type jsiiProxy_FTPServerFileDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileDeleteFailed() FTPServerFileDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileDeleteFailed_Override(f FTPServerFileDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Delete Failed.
// Experimental.
func FTPServerFileDeleteFailed_EventPattern(options *FTPServerFileDeleteFailed_FTPServerFileDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileDeleteFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileDeleteFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

