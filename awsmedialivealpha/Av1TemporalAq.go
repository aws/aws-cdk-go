package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 temporal adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1TemporalAq := medialive_alpha.Av1TemporalAq_Of(jsii.String("value"))
//
// Experimental.
type Av1TemporalAq interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1TemporalAq
type jsiiProxy_Av1TemporalAq struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1TemporalAq) Value() *string {
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
func Av1TemporalAq_Of(value *string) Av1TemporalAq {
	_init_.Initialize()

	if err := validateAv1TemporalAq_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1TemporalAq

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1TemporalAq",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1TemporalAq_DISABLED() Av1TemporalAq {
	_init_.Initialize()
	var returns Av1TemporalAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1TemporalAq",
		"DISABLED",
		&returns,
	)
	return returns
}

func Av1TemporalAq_ENABLED() Av1TemporalAq {
	_init_.Initialize()
	var returns Av1TemporalAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1TemporalAq",
		"ENABLED",
		&returns,
	)
	return returns
}

