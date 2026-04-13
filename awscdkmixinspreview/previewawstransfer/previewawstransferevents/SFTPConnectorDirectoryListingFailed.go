package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorDirectoryListingFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingFailed := awscdkmixinspreview.Events.NewSFTPConnectorDirectoryListingFailed()
//
// Experimental.
type SFTPConnectorDirectoryListingFailed interface {
}

// The jsii proxy struct for SFTPConnectorDirectoryListingFailed
type jsiiProxy_SFTPConnectorDirectoryListingFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorDirectoryListingFailed() SFTPConnectorDirectoryListingFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorDirectoryListingFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorDirectoryListingFailed_Override(s SFTPConnectorDirectoryListingFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Directory Listing Failed.
// Experimental.
func SFTPConnectorDirectoryListingFailed_EventPattern(options *SFTPConnectorDirectoryListingFailed_SFTPConnectorDirectoryListingFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorDirectoryListingFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

