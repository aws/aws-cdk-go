package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2PayloadSendFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2PayloadSendFailed := awscdkmixinspreview.Events.NewAS2PayloadSendFailed()
//
// Experimental.
type AS2PayloadSendFailed interface {
}

// The jsii proxy struct for AS2PayloadSendFailed
type jsiiProxy_AS2PayloadSendFailed struct {
	_ byte // padding
}

// Experimental.
func NewAS2PayloadSendFailed() AS2PayloadSendFailed {
	_init_.Initialize()

	j := jsiiProxy_AS2PayloadSendFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2PayloadSendFailed_Override(a AS2PayloadSendFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendFailed",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 Payload Send Failed.
// Experimental.
func AS2PayloadSendFailed_As2PayloadSendFailedPattern(options *AS2PayloadSendFailed_AS2PayloadSendFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2PayloadSendFailed_As2PayloadSendFailedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendFailed",
		"as2PayloadSendFailedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

