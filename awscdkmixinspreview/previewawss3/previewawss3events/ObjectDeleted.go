package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectDeleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectDeleted := awscdkmixinspreview.Events.NewObjectDeleted()
//
// Experimental.
type ObjectDeleted interface {
}

// The jsii proxy struct for ObjectDeleted
type jsiiProxy_ObjectDeleted struct {
	_ byte // padding
}

// Experimental.
func NewObjectDeleted() ObjectDeleted {
	_init_.Initialize()

	j := jsiiProxy_ObjectDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectDeleted_Override(o ObjectDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Deleted.
// Experimental.
func ObjectDeleted_EventPattern(options *ObjectDeleted_ObjectDeletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectDeleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

