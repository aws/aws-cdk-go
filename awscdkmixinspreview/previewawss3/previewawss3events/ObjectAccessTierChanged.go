package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.s3@ObjectAccessTierChanged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   objectAccessTierChanged := awscdkmixinspreview.Events.NewObjectAccessTierChanged()
//
// Experimental.
type ObjectAccessTierChanged interface {
}

// The jsii proxy struct for ObjectAccessTierChanged
type jsiiProxy_ObjectAccessTierChanged struct {
	_ byte // padding
}

// Experimental.
func NewObjectAccessTierChanged() ObjectAccessTierChanged {
	_init_.Initialize()

	j := jsiiProxy_ObjectAccessTierChanged{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewObjectAccessTierChanged_Override(o ObjectAccessTierChanged) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged",
		nil, // no parameters
		o,
	)
}

// EventBridge event pattern for Object Access Tier Changed.
// Experimental.
func ObjectAccessTierChanged_ObjectAccessTierChangedPattern(options *ObjectAccessTierChanged_ObjectAccessTierChangedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateObjectAccessTierChanged_ObjectAccessTierChangedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged",
		"objectAccessTierChangedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

