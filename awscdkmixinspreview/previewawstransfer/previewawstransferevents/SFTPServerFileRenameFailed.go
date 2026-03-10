package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileRenameFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileRenameFailed := awscdkmixinspreview.Events.NewSFTPServerFileRenameFailed()
//
// Experimental.
type SFTPServerFileRenameFailed interface {
}

// The jsii proxy struct for SFTPServerFileRenameFailed
type jsiiProxy_SFTPServerFileRenameFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileRenameFailed() SFTPServerFileRenameFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileRenameFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileRenameFailed_Override(s SFTPServerFileRenameFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Rename Failed.
// Experimental.
func SFTPServerFileRenameFailed_SftPServerFileRenameFailedPattern(options *SFTPServerFileRenameFailed_SFTPServerFileRenameFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileRenameFailed_SftPServerFileRenameFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameFailed",
		"sftPServerFileRenameFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

