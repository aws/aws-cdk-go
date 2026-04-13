package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorFileRetrieveFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorFileRetrieveFailed := awscdkmixinspreview.Events.NewSFTPConnectorFileRetrieveFailed()
//
// Experimental.
type SFTPConnectorFileRetrieveFailed interface {
}

// The jsii proxy struct for SFTPConnectorFileRetrieveFailed
type jsiiProxy_SFTPConnectorFileRetrieveFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorFileRetrieveFailed() SFTPConnectorFileRetrieveFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorFileRetrieveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorFileRetrieveFailed_Override(s SFTPConnectorFileRetrieveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector File Retrieve Failed.
// Experimental.
func SFTPConnectorFileRetrieveFailed_EventPattern(options *SFTPConnectorFileRetrieveFailed_SFTPConnectorFileRetrieveFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorFileRetrieveFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorFileRetrieveFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

