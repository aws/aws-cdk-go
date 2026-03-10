package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileRenameCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileRenameCompleted := awscdkmixinspreview.Events.NewSFTPServerFileRenameCompleted()
//
// Experimental.
type SFTPServerFileRenameCompleted interface {
}

// The jsii proxy struct for SFTPServerFileRenameCompleted
type jsiiProxy_SFTPServerFileRenameCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileRenameCompleted() SFTPServerFileRenameCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileRenameCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileRenameCompleted_Override(s SFTPServerFileRenameCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Rename Completed.
// Experimental.
func SFTPServerFileRenameCompleted_SftPServerFileRenameCompletedPattern(options *SFTPServerFileRenameCompleted_SFTPServerFileRenameCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileRenameCompleted_SftPServerFileRenameCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileRenameCompleted",
		"sftPServerFileRenameCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

