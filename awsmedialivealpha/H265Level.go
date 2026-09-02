package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265Level := medialive_alpha.H265Level_H265_LEVEL_1()
//
// Experimental.
type H265Level interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265Level
type jsiiProxy_H265Level struct {
	_ byte // padding
}

func (j *jsiiProxy_H265Level) Value() *string {
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
func H265Level_Of(value *string) H265Level {
	_init_.Initialize()

	if err := validateH265Level_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265Level

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265Level_H265_LEVEL_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_2() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_2",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_2_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_2_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_3() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_3",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_3_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_3_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_4() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_4",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_4_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_4_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_5() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_5",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_5_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_5_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_5_2() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_5_2",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_6() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_6",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_6_1() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_6_1",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_6_2() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_6_2",
		&returns,
	)
	return returns
}

func H265Level_H265_LEVEL_AUTO() H265Level {
	_init_.Initialize()
	var returns H265Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Level",
		"H265_LEVEL_AUTO",
		&returns,
	)
	return returns
}

