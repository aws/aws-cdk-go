package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileDeleteCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileDeleteCompleted := awscdkmixinspreview.Events.NewSFTPServerFileDeleteCompleted()
//
// Experimental.
type SFTPServerFileDeleteCompleted interface {
}

// The jsii proxy struct for SFTPServerFileDeleteCompleted
type jsiiProxy_SFTPServerFileDeleteCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileDeleteCompleted() SFTPServerFileDeleteCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileDeleteCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileDeleteCompleted_Override(s SFTPServerFileDeleteCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Delete Completed.
// Experimental.
func SFTPServerFileDeleteCompleted_SftPServerFileDeleteCompletedPattern(options *SFTPServerFileDeleteCompleted_SFTPServerFileDeleteCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileDeleteCompleted_SftPServerFileDeleteCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDeleteCompleted",
		"sftPServerFileDeleteCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

