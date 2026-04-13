package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerDirectoryDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerDirectoryDeleteFailed := awscdkmixinspreview.Events.NewFTPServerDirectoryDeleteFailed()
//
// Experimental.
type FTPServerDirectoryDeleteFailed interface {
}

// The jsii proxy struct for FTPServerDirectoryDeleteFailed
type jsiiProxy_FTPServerDirectoryDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerDirectoryDeleteFailed() FTPServerDirectoryDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerDirectoryDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerDirectoryDeleteFailed_Override(f FTPServerDirectoryDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server Directory Delete Failed.
// Experimental.
func FTPServerDirectoryDeleteFailed_EventPattern(options *FTPServerDirectoryDeleteFailed_FTPServerDirectoryDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerDirectoryDeleteFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

