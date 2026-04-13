package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerDirectoryDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerDirectoryDeleteCompleted := awscdkmixinspreview.Events.NewFTPServerDirectoryDeleteCompleted()
//
// Experimental.
type FTPServerDirectoryDeleteCompleted interface {
}

// The jsii proxy struct for FTPServerDirectoryDeleteCompleted
type jsiiProxy_FTPServerDirectoryDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerDirectoryDeleteCompleted() FTPServerDirectoryDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPServerDirectoryDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerDirectoryDeleteCompleted_Override(f FTPServerDirectoryDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server Directory Delete Completed.
// Experimental.
func FTPServerDirectoryDeleteCompleted_EventPattern(options *FTPServerDirectoryDeleteCompleted_FTPServerDirectoryDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerDirectoryDeleteCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerDirectoryDeleteCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

