package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerDirectoryCreateFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerDirectoryCreateFailed := awscdkmixinspreview.Events.NewFTPSServerDirectoryCreateFailed()
//
// Experimental.
type FTPSServerDirectoryCreateFailed interface {
}

// The jsii proxy struct for FTPSServerDirectoryCreateFailed
type jsiiProxy_FTPSServerDirectoryCreateFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerDirectoryCreateFailed() FTPSServerDirectoryCreateFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerDirectoryCreateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerDirectoryCreateFailed_Override(f FTPSServerDirectoryCreateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server Directory Create Failed.
// Experimental.
func FTPSServerDirectoryCreateFailed_EventPattern(options *FTPSServerDirectoryCreateFailed_FTPSServerDirectoryCreateFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerDirectoryCreateFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerDirectoryCreateFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

