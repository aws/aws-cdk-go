package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectRestoreCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreCompleted := awscdkmixinspreview.Events.NewObjectRestoreCompleted()
//
// Experimental.
type ObjectRestoreCompleted interface {
}

// The jsii proxy struct for ObjectRestoreCompleted
type jsiiProxy_ObjectRestoreCompleted struct {
	_ byte // padding
}

// Experimental.
func NewObjectRestoreCompleted() ObjectRestoreCompleted {
	_init_.Initialize()

	j := jsiiProxy_ObjectRestoreCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectRestoreCompleted_Override(o ObjectRestoreCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Restore Completed.
// Experimental.
func ObjectRestoreCompleted_EventPattern(options *ObjectRestoreCompleted_ObjectRestoreCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectRestoreCompleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

