package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Nielsen ID3 passthrough behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsNielsenId3Behavior := medialive_alpha.M2tsNielsenId3Behavior_Of(jsii.String("value"))
//
// Experimental.
type M2tsNielsenId3Behavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsNielsenId3Behavior
type jsiiProxy_M2tsNielsenId3Behavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsNielsenId3Behavior) Value() *string {
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
func M2tsNielsenId3Behavior_Of(value *string) M2tsNielsenId3Behavior {
	_init_.Initialize()

	if err := validateM2tsNielsenId3Behavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsNielsenId3Behavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsNielsenId3Behavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsNielsenId3Behavior_NO_PASSTHROUGH() M2tsNielsenId3Behavior {
	_init_.Initialize()
	var returns M2tsNielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsNielsenId3Behavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func M2tsNielsenId3Behavior_PASSTHROUGH() M2tsNielsenId3Behavior {
	_init_.Initialize()
	var returns M2tsNielsenId3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsNielsenId3Behavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

