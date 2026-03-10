package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPServerFileDownloadFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPServerFileDownloadFailed := awscdkmixinspreview.Events.NewSFTPServerFileDownloadFailed()
//
// Experimental.
type SFTPServerFileDownloadFailed interface {
}

// The jsii proxy struct for SFTPServerFileDownloadFailed
type jsiiProxy_SFTPServerFileDownloadFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPServerFileDownloadFailed() SFTPServerFileDownloadFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPServerFileDownloadFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPServerFileDownloadFailed_Override(s SFTPServerFileDownloadFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Server File Download Failed.
// Experimental.
func SFTPServerFileDownloadFailed_SftPServerFileDownloadFailedPattern(options *SFTPServerFileDownloadFailed_SFTPServerFileDownloadFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPServerFileDownloadFailed_SftPServerFileDownloadFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPServerFileDownloadFailed",
		"sftPServerFileDownloadFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

