package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2PayloadSendCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2PayloadSendCompleted := awscdkmixinspreview.Events.NewAS2PayloadSendCompleted()
//
// Experimental.
type AS2PayloadSendCompleted interface {
}

// The jsii proxy struct for AS2PayloadSendCompleted
type jsiiProxy_AS2PayloadSendCompleted struct {
	_ byte // padding
}

// Experimental.
func NewAS2PayloadSendCompleted() AS2PayloadSendCompleted {
	_init_.Initialize()

	j := jsiiProxy_AS2PayloadSendCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2PayloadSendCompleted_Override(a AS2PayloadSendCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendCompleted",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 Payload Send Completed.
// Experimental.
func AS2PayloadSendCompleted_EventPattern(options *AS2PayloadSendCompleted_AS2PayloadSendCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2PayloadSendCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadSendCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

