package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectTagsAdded.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectTagsAdded := awscdkmixinspreview.Events.NewObjectTagsAdded()
//
// Experimental.
type ObjectTagsAdded interface {
}

// The jsii proxy struct for ObjectTagsAdded
type jsiiProxy_ObjectTagsAdded struct {
	_ byte // padding
}

// Experimental.
func NewObjectTagsAdded() ObjectTagsAdded {
	_init_.Initialize()

	j := jsiiProxy_ObjectTagsAdded{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectTagsAdded_Override(o ObjectTagsAdded) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Tags Added.
// Experimental.
func ObjectTagsAdded_EventPattern(options *ObjectTagsAdded_ObjectTagsAddedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectTagsAdded_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

