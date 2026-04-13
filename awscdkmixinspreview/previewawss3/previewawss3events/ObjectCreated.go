package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectCreated := awscdkmixinspreview.Events.NewObjectCreated()
//
// Experimental.
type ObjectCreated interface {
}

// The jsii proxy struct for ObjectCreated
type jsiiProxy_ObjectCreated struct {
	_ byte // padding
}

// Experimental.
func NewObjectCreated() ObjectCreated {
	_init_.Initialize()

	j := jsiiProxy_ObjectCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectCreated_Override(o ObjectCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Created.
// Experimental.
func ObjectCreated_EventPattern(options *ObjectCreated_ObjectCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectCreated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

