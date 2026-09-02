package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1Level := medialive_alpha.Av1Level_AV1_LEVEL_2()
//
// Experimental.
type Av1Level interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1Level
type jsiiProxy_Av1Level struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1Level) Value() *string {
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
func Av1Level_Of(value *string) Av1Level {
	_init_.Initialize()

	if err := validateAv1Level_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1Level

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1Level_AV1_LEVEL_2() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_2",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_2_1() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_2_1",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_3() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_3",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_3_1() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_3_1",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_4() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_4",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_4_1() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_4_1",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_5() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_5",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_5_1() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_5_1",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_5_2() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_5_2",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_5_3() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_5_3",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_6() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_6",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_6_1() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_6_1",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_6_2() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_6_2",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_6_3() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_6_3",
		&returns,
	)
	return returns
}

func Av1Level_AV1_LEVEL_AUTO() Av1Level {
	_init_.Initialize()
	var returns Av1Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1Level",
		"AV1_LEVEL_AUTO",
		&returns,
	)
	return returns
}

