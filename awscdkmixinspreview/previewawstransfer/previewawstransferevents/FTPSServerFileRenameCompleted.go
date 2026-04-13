package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileRenameCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileRenameCompleted := awscdkmixinspreview.Events.NewFTPSServerFileRenameCompleted()
//
// Experimental.
type FTPSServerFileRenameCompleted interface {
}

// The jsii proxy struct for FTPSServerFileRenameCompleted
type jsiiProxy_FTPSServerFileRenameCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileRenameCompleted() FTPSServerFileRenameCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileRenameCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileRenameCompleted_Override(f FTPSServerFileRenameCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Rename Completed.
// Experimental.
func FTPSServerFileRenameCompleted_EventPattern(options *FTPSServerFileRenameCompleted_FTPSServerFileRenameCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileRenameCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

