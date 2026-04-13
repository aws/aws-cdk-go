package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerDirectoryDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerDirectoryDeleteFailed := awscdkmixinspreview.Events.NewFTPSServerDirectoryDeleteFailed()
//
// Experimental.
type FTPSServerDirectoryDeleteFailed interface {
}

// The jsii proxy struct for FTPSServerDirectoryDeleteFailed
type jsiiProxy_FTPSServerDirectoryDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerDirectoryDeleteFailed() FTPSServerDirectoryDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerDirectoryDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerDirectoryDeleteFailed_Override(f FTPSServerDirectoryDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server Directory Delete Failed.
// Experimental.
func FTPSServerDirectoryDeleteFailed_EventPattern(options *FTPSServerDirectoryDeleteFailed_FTPSServerDirectoryDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerDirectoryDeleteFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryDeleteFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

