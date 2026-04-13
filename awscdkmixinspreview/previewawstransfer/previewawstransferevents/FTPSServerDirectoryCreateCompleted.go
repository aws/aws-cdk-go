package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerDirectoryCreateCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerDirectoryCreateCompleted := awscdkmixinspreview.Events.NewFTPSServerDirectoryCreateCompleted()
//
// Experimental.
type FTPSServerDirectoryCreateCompleted interface {
}

// The jsii proxy struct for FTPSServerDirectoryCreateCompleted
type jsiiProxy_FTPSServerDirectoryCreateCompleted struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerDirectoryCreateCompleted() FTPSServerDirectoryCreateCompleted {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerDirectoryCreateCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerDirectoryCreateCompleted_Override(f FTPSServerDirectoryCreateCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateCompleted",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server Directory Create Completed.
// Experimental.
func FTPSServerDirectoryCreateCompleted_EventPattern(options *FTPSServerDirectoryCreateCompleted_FTPSServerDirectoryCreateCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerDirectoryCreateCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

