package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Motion graphics insertion state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   motionGraphicsInsertion := medialive_alpha.MotionGraphicsInsertion_Of(jsii.String("value"))
//
// Experimental.
type MotionGraphicsInsertion interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MotionGraphicsInsertion
type jsiiProxy_MotionGraphicsInsertion struct {
	_ byte // padding
}

func (j *jsiiProxy_MotionGraphicsInsertion) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func MotionGraphicsInsertion_Of(value *string) MotionGraphicsInsertion {
	_init_.Initialize()

	if err := validateMotionGraphicsInsertion_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MotionGraphicsInsertion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MotionGraphicsInsertion",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MotionGraphicsInsertion_DISABLED() MotionGraphicsInsertion {
	_init_.Initialize()
	var returns MotionGraphicsInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MotionGraphicsInsertion",
		"DISABLED",
		&returns,
	)
	return returns
}

func MotionGraphicsInsertion_ENABLED() MotionGraphicsInsertion {
	_init_.Initialize()
	var returns MotionGraphicsInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MotionGraphicsInsertion",
		"ENABLED",
		&returns,
	)
	return returns
}

