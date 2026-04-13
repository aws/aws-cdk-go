package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorFileRetrieveCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileRetrieveCompleted := awscdkmixinspreview.Events.NewSFTPConnectorFileRetrieveCompleted()
//
// Experimental.
type SFTPConnectorFileRetrieveCompleted interface {
}

// The jsii proxy struct for SFTPConnectorFileRetrieveCompleted
type jsiiProxy_SFTPConnectorFileRetrieveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorFileRetrieveCompleted() SFTPConnectorFileRetrieveCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorFileRetrieveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorFileRetrieveCompleted_Override(s SFTPConnectorFileRetrieveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector File Retrieve Completed.
// Experimental.
func SFTPConnectorFileRetrieveCompleted_EventPattern(options *SFTPConnectorFileRetrieveCompleted_SFTPConnectorFileRetrieveCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorFileRetrieveCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

