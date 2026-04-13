package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerDirectoryDeleteFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerDirectoryDeleteFailed := awscdkmixinspreview.Events.NewSFTPServerDirectoryDeleteFailed()
//
// Experimental.
type SFTPServerDirectoryDeleteFailed interface {
}

// The jsii proxy struct for SFTPServerDirectoryDeleteFailed
type jsiiProxy_SFTPServerDirectoryDeleteFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerDirectoryDeleteFailed() SFTPServerDirectoryDeleteFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerDirectoryDeleteFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerDirectoryDeleteFailed_Override(s SFTPServerDirectoryDeleteFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server Directory Delete Failed.
// Experimental.
func SFTPServerDirectoryDeleteFailed_EventPattern(options *SFTPServerDirectoryDeleteFailed_SFTPServerDirectoryDeleteFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerDirectoryDeleteFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

