package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectACLUpdated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectACLUpdated := awscdkmixinspreview.Events.NewObjectACLUpdated()
//
// Experimental.
type ObjectACLUpdated interface {
}

// The jsii proxy struct for ObjectACLUpdated
type jsiiProxy_ObjectACLUpdated struct {
	_ byte // padding
}

// Experimental.
func NewObjectACLUpdated() ObjectACLUpdated {
	_init_.Initialize()

	j := jsiiProxy_ObjectACLUpdated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectACLUpdated_Override(o ObjectACLUpdated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object ACL Updated.
// Experimental.
func ObjectACLUpdated_EventPattern(options *ObjectACLUpdated_ObjectACLUpdatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectACLUpdated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

