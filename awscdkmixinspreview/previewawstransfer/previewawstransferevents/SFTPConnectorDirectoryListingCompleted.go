package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorDirectoryListingCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorDirectoryListingCompleted := awscdkmixinspreview.Events.NewSFTPConnectorDirectoryListingCompleted()
//
// Experimental.
type SFTPConnectorDirectoryListingCompleted interface {
}

// The jsii proxy struct for SFTPConnectorDirectoryListingCompleted
type jsiiProxy_SFTPConnectorDirectoryListingCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorDirectoryListingCompleted() SFTPConnectorDirectoryListingCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorDirectoryListingCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorDirectoryListingCompleted_Override(s SFTPConnectorDirectoryListingCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Directory Listing Completed.
// Experimental.
func SFTPConnectorDirectoryListingCompleted_EventPattern(options *SFTPConnectorDirectoryListingCompleted_SFTPConnectorDirectoryListingCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorDirectoryListingCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorDirectoryListingCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

