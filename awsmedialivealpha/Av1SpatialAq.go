package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 spatial adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1SpatialAq := medialive_alpha.Av1SpatialAq_Of(jsii.String("value"))
//
// Experimental.
type Av1SpatialAq interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1SpatialAq
type jsiiProxy_Av1SpatialAq struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1SpatialAq) Value() *string {
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
func Av1SpatialAq_Of(value *string) Av1SpatialAq {
	_init_.Initialize()

	if err := validateAv1SpatialAq_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1SpatialAq

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1SpatialAq",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1SpatialAq_DISABLED() Av1SpatialAq {
	_init_.Initialize()
	var returns Av1SpatialAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1SpatialAq",
		"DISABLED",
		&returns,
	)
	return returns
}

func Av1SpatialAq_ENABLED() Av1SpatialAq {
	_init_.Initialize()
	var returns Av1SpatialAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1SpatialAq",
		"ENABLED",
		&returns,
	)
	return returns
}

