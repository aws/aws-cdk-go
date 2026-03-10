package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@FTPSServerFileRenameFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fTPSServerFileRenameFailed := awscdkmixinspreview.Events.NewFTPSServerFileRenameFailed()
//
// Experimental.
type FTPSServerFileRenameFailed interface {
}

// The jsii proxy struct for FTPSServerFileRenameFailed
type jsiiProxy_FTPSServerFileRenameFailed struct {
	_ byte // padding
}

// Experimental.
func NewFTPSServerFileRenameFailed() FTPSServerFileRenameFailed {
	_init_.Initialize()

	j := jsiiProxy_FTPSServerFileRenameFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFTPSServerFileRenameFailed_Override(f FTPSServerFileRenameFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameFailed",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FTPS Server File Rename Failed.
// Experimental.
func FTPSServerFileRenameFailed_FtpSServerFileRenameFailedPattern(options *FTPSServerFileRenameFailed_FTPSServerFileRenameFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFTPSServerFileRenameFailed_FtpSServerFileRenameFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.FTPSServerFileRenameFailed",
		"ftpSServerFileRenameFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

