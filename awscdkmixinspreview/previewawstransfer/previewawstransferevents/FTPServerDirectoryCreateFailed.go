package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerDirectoryCreateFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerDirectoryCreateFailed := awscdkmixinspreview.Events.NewFTPServerDirectoryCreateFailed()
//
// Experimental.
type FTPServerDirectoryCreateFailed interface {
}

// The jsii proxy struct for FTPServerDirectoryCreateFailed
type jsiiProxy_FTPServerDirectoryCreateFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerDirectoryCreateFailed() FTPServerDirectoryCreateFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerDirectoryCreateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerDirectoryCreateFailed_Override(f FTPServerDirectoryCreateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server Directory Create Failed.
// Experimental.
func FTPServerDirectoryCreateFailed_EventPattern(options *FTPServerDirectoryCreateFailed_FTPServerDirectoryCreateFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerDirectoryCreateFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

