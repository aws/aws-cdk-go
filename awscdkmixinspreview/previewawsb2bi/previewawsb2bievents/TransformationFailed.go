package previewawsb2bievents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.b2bi@TransformationFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformationFailed := awscdkmixinspreview.Events.NewTransformationFailed()
//
// Experimental.
type TransformationFailed interface {
}

// The jsii proxy struct for TransformationFailed
type jsiiProxy_TransformationFailed struct {
	_ byte // padding
}

// Experimental.
func NewTransformationFailed() TransformationFailed {
	_init_.Initialize()

	j := jsiiProxy_TransformationFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTransformationFailed_Override(t TransformationFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Transformation Failed.
// Experimental.
func TransformationFailed_EventPattern(options *TransformationFailed_TransformationFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTransformationFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

