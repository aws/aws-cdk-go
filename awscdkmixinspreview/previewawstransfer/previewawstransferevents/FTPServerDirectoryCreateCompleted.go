package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerDirectoryCreateCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerDirectoryCreateCompleted := awscdkmixinspreview.Events.NewFTPServerDirectoryCreateCompleted()
//
// Experimental.
type FTPServerDirectoryCreateCompleted interface {
}

// The jsii proxy struct for FTPServerDirectoryCreateCompleted
type jsiiProxy_FTPServerDirectoryCreateCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerDirectoryCreateCompleted() FTPServerDirectoryCreateCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerDirectoryCreateCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerDirectoryCreateCompleted_Override(f FTPServerDirectoryCreateCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server Directory Create Completed.
// Experimental.
func FTPServerDirectoryCreateCompleted_EventPattern(options *FTPServerDirectoryCreateCompleted_FTPServerDirectoryCreateCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerDirectoryCreateCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryCreateCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

