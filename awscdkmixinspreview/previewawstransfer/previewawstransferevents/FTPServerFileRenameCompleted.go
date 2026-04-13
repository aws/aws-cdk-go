package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileRenameCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileRenameCompleted := awscdkmixinspreview.Events.NewFTPServerFileRenameCompleted()
//
// Experimental.
type FTPServerFileRenameCompleted interface {
}

// The jsii proxy struct for FTPServerFileRenameCompleted
type jsiiProxy_FTPServerFileRenameCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileRenameCompleted() FTPServerFileRenameCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileRenameCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileRenameCompleted_Override(f FTPServerFileRenameCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Rename Completed.
// Experimental.
func FTPServerFileRenameCompleted_EventPattern(options *FTPServerFileRenameCompleted_FTPServerFileRenameCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileRenameCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

