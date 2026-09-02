package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Video scaling behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   scalingBehavior := medialive_alpha.ScalingBehavior_DEFAULT()
//
// Experimental.
type ScalingBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ScalingBehavior
type jsiiProxy_ScalingBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_ScalingBehavior) Value() *string {
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
func ScalingBehavior_Of(value *string) ScalingBehavior {
	_init_.Initialize()

	if err := validateScalingBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ScalingBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ScalingBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ScalingBehavior_DEFAULT() ScalingBehavior {
	_init_.Initialize()
	var returns ScalingBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ScalingBehavior",
		"DEFAULT",
		&returns,
	)
	return returns
}

func ScalingBehavior_SMART_CROP() ScalingBehavior {
	_init_.Initialize()
	var returns ScalingBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ScalingBehavior",
		"SMART_CROP",
		&returns,
	)
	return returns
}

func ScalingBehavior_STRETCH_TO_OUTPUT() ScalingBehavior {
	_init_.Initialize()
	var returns ScalingBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ScalingBehavior",
		"STRETCH_TO_OUTPUT",
		&returns,
	)
	return returns
}

