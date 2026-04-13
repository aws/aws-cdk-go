package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorRemoteMoveFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteMoveFailed := awscdkmixinspreview.Events.NewSFTPConnectorRemoteMoveFailed()
//
// Experimental.
type SFTPConnectorRemoteMoveFailed interface {
}

// The jsii proxy struct for SFTPConnectorRemoteMoveFailed
type jsiiProxy_SFTPConnectorRemoteMoveFailed struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorRemoteMoveFailed() SFTPConnectorRemoteMoveFailed {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorRemoteMoveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorRemoteMoveFailed_Override(s SFTPConnectorRemoteMoveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveFailed",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Remote Move Failed.
// Experimental.
func SFTPConnectorRemoteMoveFailed_EventPattern(options *SFTPConnectorRemoteMoveFailed_SFTPConnectorRemoteMoveFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorRemoteMoveFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

