package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Temporal filter strength.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   temporalFilterStrength := medialive_alpha.TemporalFilterStrength_AUTO()
//
// Experimental.
type TemporalFilterStrength interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TemporalFilterStrength
type jsiiProxy_TemporalFilterStrength struct {
	_ byte // padding
}

func (j *jsiiProxy_TemporalFilterStrength) Value() *string {
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
func TemporalFilterStrength_Of(value *string) TemporalFilterStrength {
	_init_.Initialize()

	if err := validateTemporalFilterStrength_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TemporalFilterStrength

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TemporalFilterStrength_AUTO() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"AUTO",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_1() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_1",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_10() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_10",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_11() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_11",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_12() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_12",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_13() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_13",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_14() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_14",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_15() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_15",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_16() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_16",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_2() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_2",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_3() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_3",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_4() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_4",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_5() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_5",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_6() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_6",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_7() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_7",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_8() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_8",
		&returns,
	)
	return returns
}

func TemporalFilterStrength_STRENGTH_9() TemporalFilterStrength {
	_init_.Initialize()
	var returns TemporalFilterStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TemporalFilterStrength",
		"STRENGTH_9",
		&returns,
	)
	return returns
}

