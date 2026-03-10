package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerDirectoryDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerDirectoryDeleteCompleted := awscdkmixinspreview.Events.NewSFTPServerDirectoryDeleteCompleted()
//
// Experimental.
type SFTPServerDirectoryDeleteCompleted interface {
}

// The jsii proxy struct for SFTPServerDirectoryDeleteCompleted
type jsiiProxy_SFTPServerDirectoryDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerDirectoryDeleteCompleted() SFTPServerDirectoryDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerDirectoryDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerDirectoryDeleteCompleted_Override(s SFTPServerDirectoryDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server Directory Delete Completed.
// Experimental.
func SFTPServerDirectoryDeleteCompleted_SftPServerDirectoryDeleteCompletedPattern(options *SFTPServerDirectoryDeleteCompleted_SFTPServerDirectoryDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerDirectoryDeleteCompleted_SftPServerDirectoryDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerDirectoryDeleteCompleted",
		"sftPServerDirectoryDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

