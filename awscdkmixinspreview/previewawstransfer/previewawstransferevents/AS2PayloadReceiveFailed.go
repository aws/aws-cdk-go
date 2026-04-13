package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2PayloadReceiveFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2PayloadReceiveFailed := awscdkmixinspreview.Events.NewAS2PayloadReceiveFailed()
//
// Experimental.
type AS2PayloadReceiveFailed interface {
}

// The jsii proxy struct for AS2PayloadReceiveFailed
type jsiiProxy_AS2PayloadReceiveFailed struct {
	_ byte // padding
}

// Experimental.
func NewAS2PayloadReceiveFailed() AS2PayloadReceiveFailed {
	_init_.Initialize()

	j := jsiiProxy_AS2PayloadReceiveFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2PayloadReceiveFailed_Override(a AS2PayloadReceiveFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveFailed",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 Payload Receive Failed.
// Experimental.
func AS2PayloadReceiveFailed_EventPattern(options *AS2PayloadReceiveFailed_AS2PayloadReceiveFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2PayloadReceiveFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

