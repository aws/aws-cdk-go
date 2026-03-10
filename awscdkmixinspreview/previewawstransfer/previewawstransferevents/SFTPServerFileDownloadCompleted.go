package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileDownloadCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileDownloadCompleted := awscdkmixinspreview.Events.NewSFTPServerFileDownloadCompleted()
//
// Experimental.
type SFTPServerFileDownloadCompleted interface {
}

// The jsii proxy struct for SFTPServerFileDownloadCompleted
type jsiiProxy_SFTPServerFileDownloadCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileDownloadCompleted() SFTPServerFileDownloadCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileDownloadCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileDownloadCompleted_Override(s SFTPServerFileDownloadCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Download Completed.
// Experimental.
func SFTPServerFileDownloadCompleted_SftPServerFileDownloadCompletedPattern(options *SFTPServerFileDownloadCompleted_SFTPServerFileDownloadCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileDownloadCompleted_SftPServerFileDownloadCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadCompleted",
		"sftPServerFileDownloadCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

