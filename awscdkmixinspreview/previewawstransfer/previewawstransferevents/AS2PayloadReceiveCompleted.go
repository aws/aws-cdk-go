package previewawstransferevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.transfer@AS2PayloadReceiveCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aS2PayloadReceiveCompleted := awscdkmixinspreview.Events.NewAS2PayloadReceiveCompleted()
//
// Experimental.
type AS2PayloadReceiveCompleted interface {
}

// The jsii proxy struct for AS2PayloadReceiveCompleted
type jsiiProxy_AS2PayloadReceiveCompleted struct {
	_ byte // padding
}

// Experimental.
func NewAS2PayloadReceiveCompleted() AS2PayloadReceiveCompleted {
	_init_.Initialize()

	j := jsiiProxy_AS2PayloadReceiveCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAS2PayloadReceiveCompleted_Override(a AS2PayloadReceiveCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveCompleted",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for AS2 Payload Receive Completed.
// Experimental.
func AS2PayloadReceiveCompleted_As2PayloadReceiveCompletedPattern(options *AS2PayloadReceiveCompleted_AS2PayloadReceiveCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAS2PayloadReceiveCompleted_As2PayloadReceiveCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.events.AS2PayloadReceiveCompleted",
		"as2PayloadReceiveCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

