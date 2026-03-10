package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2MDNReceiveCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2MDNReceiveCompleted := awscdkmixinspreview.Events.NewAS2MDNReceiveCompleted()
//
// Experimental.
type AS2MDNReceiveCompleted interface {
}

// The jsii proxy struct for AS2MDNReceiveCompleted
type jsiiProxy_AS2MDNReceiveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewAS2MDNReceiveCompleted() AS2MDNReceiveCompleted {
	_init_.Initialize()

	j := jsiiProxy_AS2MDNReceiveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2MDNReceiveCompleted_Override(a AS2MDNReceiveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveCompleted",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 MDN Receive Completed.
// Experimental.
func AS2MDNReceiveCompleted_As2MDNReceiveCompletedPattern(options *AS2MDNReceiveCompleted_AS2MDNReceiveCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2MDNReceiveCompleted_As2MDNReceiveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2MDNReceiveCompleted",
		"as2MDNReceiveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

