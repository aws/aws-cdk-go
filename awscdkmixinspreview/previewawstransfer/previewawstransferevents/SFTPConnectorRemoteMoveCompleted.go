package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@SFTPConnectorRemoteMoveCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sFTPConnectorRemoteMoveCompleted := awscdkmixinspreview.Events.NewSFTPConnectorRemoteMoveCompleted()
//
// Experimental.
type SFTPConnectorRemoteMoveCompleted interface {
}

// The jsii proxy struct for SFTPConnectorRemoteMoveCompleted
type jsiiProxy_SFTPConnectorRemoteMoveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewSFTPConnectorRemoteMoveCompleted() SFTPConnectorRemoteMoveCompleted {
	_init_.Initialize()

	j := jsiiProxy_SFTPConnectorRemoteMoveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSFTPConnectorRemoteMoveCompleted_Override(s SFTPConnectorRemoteMoveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveCompleted",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for SFTP Connector Remote Move Completed.
// Experimental.
func SFTPConnectorRemoteMoveCompleted_SftPConnectorRemoteMoveCompletedPattern(options *SFTPConnectorRemoteMoveCompleted_SFTPConnectorRemoteMoveCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSFTPConnectorRemoteMoveCompleted_SftPConnectorRemoteMoveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.SFTPConnectorRemoteMoveCompleted",
		"sftPConnectorRemoteMoveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

