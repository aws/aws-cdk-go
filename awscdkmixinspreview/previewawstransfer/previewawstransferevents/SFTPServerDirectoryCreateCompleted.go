package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerDirectoryCreateCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerDirectoryCreateCompleted := awscdkmixinspreview.Events.NewSFTPServerDirectoryCreateCompleted()
//
// Experimental.
type SFTPServerDirectoryCreateCompleted interface {
}

// The jsii proxy struct for SFTPServerDirectoryCreateCompleted
type jsiiProxy_SFTPServerDirectoryCreateCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerDirectoryCreateCompleted() SFTPServerDirectoryCreateCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerDirectoryCreateCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerDirectoryCreateCompleted_Override(s SFTPServerDirectoryCreateCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server Directory Create Completed.
// Experimental.
func SFTPServerDirectoryCreateCompleted_EventPattern(options *SFTPServerDirectoryCreateCompleted_SFTPServerDirectoryCreateCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerDirectoryCreateCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryCreateCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

