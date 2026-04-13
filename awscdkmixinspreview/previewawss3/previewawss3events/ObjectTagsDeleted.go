package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectTagsDeleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectTagsDeleted := awscdkmixinspreview.Events.NewObjectTagsDeleted()
//
// Experimental.
type ObjectTagsDeleted interface {
}

// The jsii proxy struct for ObjectTagsDeleted
type jsiiProxy_ObjectTagsDeleted struct {
	_ byte // padding
}

// Experimental.
func NewObjectTagsDeleted() ObjectTagsDeleted {
	_init_.Initialize()

	j := jsiiProxy_ObjectTagsDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectTagsDeleted_Override(o ObjectTagsDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Tags Deleted.
// Experimental.
func ObjectTagsDeleted_EventPattern(options *ObjectTagsDeleted_ObjectTagsDeletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectTagsDeleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

