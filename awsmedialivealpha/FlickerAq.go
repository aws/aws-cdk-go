package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Flicker adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   flickerAq := medialive_alpha.FlickerAq_Of(jsii.String("value"))
//
// Experimental.
type FlickerAq interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for FlickerAq
type jsiiProxy_FlickerAq struct {
	_ byte // padding
}

func (j *jsiiProxy_FlickerAq) Value() *string {
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
func FlickerAq_Of(value *string) FlickerAq {
	_init_.Initialize()

	if err := validateFlickerAq_OfParameters(value); err != nil {
		panic(err)
	}
	var returns FlickerAq

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FlickerAq",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FlickerAq_DISABLED() FlickerAq {
	_init_.Initialize()
	var returns FlickerAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FlickerAq",
		"DISABLED",
		&returns,
	)
	return returns
}

func FlickerAq_ENABLED() FlickerAq {
	_init_.Initialize()
	var returns FlickerAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FlickerAq",
		"ENABLED",
		&returns,
	)
	return returns
}

