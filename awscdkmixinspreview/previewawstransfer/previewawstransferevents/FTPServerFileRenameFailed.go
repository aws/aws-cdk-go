package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPServerFileRenameFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPServerFileRenameFailed := awscdkmixinspreview.Events.NewFTPServerFileRenameFailed()
//
// Experimental.
type FTPServerFileRenameFailed interface {
}

// The jsii proxy struct for FTPServerFileRenameFailed
type jsiiProxy_FTPServerFileRenameFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPServerFileRenameFailed() FTPServerFileRenameFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPServerFileRenameFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPServerFileRenameFailed_Override(f FTPServerFileRenameFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTP Server File Rename Failed.
// Experimental.
func FTPServerFileRenameFailed_FtpServerFileRenameFailedPattern(options *FTPServerFileRenameFailed_FTPServerFileRenameFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPServerFileRenameFailed_FtpServerFileRenameFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPServerFileRenameFailed",
		"ftpServerFileRenameFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

