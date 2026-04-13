package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2MDNSendCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNSendCompleted := awscdkmixinspreview.Events.NewAS2MDNSendCompleted()
//
// Experimental.
type AS2MDNSendCompleted interface {
}

// The jsii proxy struct for AS2MDNSendCompleted
type jsiiProxy_AS2MDNSendCompleted struct {
	_ byte // padding
}

// Experimental.
func NewAS2MDNSendCompleted() AS2MDNSendCompleted {
	_init_.Initialize()

	j := jsiiProxy_AS2MDNSendCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2MDNSendCompleted_Override(a AS2MDNSendCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendCompleted",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 MDN Send Completed.
// Experimental.
func AS2MDNSendCompleted_EventPattern(options *AS2MDNSendCompleted_AS2MDNSendCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2MDNSendCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNSendCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

