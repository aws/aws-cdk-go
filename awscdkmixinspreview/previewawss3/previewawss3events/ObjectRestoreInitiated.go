package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectRestoreInitiated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectRestoreInitiated := awscdkmixinspreview.Events.NewObjectRestoreInitiated()
//
// Experimental.
type ObjectRestoreInitiated interface {
}

// The jsii proxy struct for ObjectRestoreInitiated
type jsiiProxy_ObjectRestoreInitiated struct {
	_ byte // padding
}

// Experimental.
func NewObjectRestoreInitiated() ObjectRestoreInitiated {
	_init_.Initialize()

	j := jsiiProxy_ObjectRestoreInitiated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectRestoreInitiated_Override(o ObjectRestoreInitiated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Restore Initiated.
// Experimental.
func ObjectRestoreInitiated_EventPattern(options *ObjectRestoreInitiated_ObjectRestoreInitiatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectRestoreInitiated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

