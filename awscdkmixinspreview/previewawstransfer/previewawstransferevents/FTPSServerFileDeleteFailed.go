package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileDeleteFailed := awscdkmixinspreview.Events.NewFTPSServerFileDeleteFailed()
//
// Experimental.
type FTPSServerFileDeleteFailed interface {
}

// The jsii proxy struct for FTPSServerFileDeleteFailed
type jsiiProxy_FTPSServerFileDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileDeleteFailed() FTPSServerFileDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileDeleteFailed_Override(f FTPSServerFileDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Delete Failed.
// Experimental.
func FTPSServerFileDeleteFailed_FtpSServerFileDeleteFailedPattern(options *FTPSServerFileDeleteFailed_FTPSServerFileDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileDeleteFailed_FtpSServerFileDeleteFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileDeleteFailed",
		"ftpSServerFileDeleteFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

