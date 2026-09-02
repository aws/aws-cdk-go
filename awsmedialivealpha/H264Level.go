package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 level.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264Level := medialive_alpha.H264Level_H264_LEVEL_1()
//
// Experimental.
type H264Level interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264Level
type jsiiProxy_H264Level struct {
	_ byte // padding
}

func (j *jsiiProxy_H264Level) Value() *string {
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
func H264Level_Of(value *string) H264Level {
	_init_.Initialize()

	if err := validateH264Level_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264Level

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264Level_H264_LEVEL_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_1_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_1_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_1_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_1_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_1_3() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_1_3",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_2_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_2_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_2_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_2_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_3() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_3",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_3_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_3_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_3_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_3_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_4() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_4",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_4_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_4_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_4_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_4_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_5() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_5",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_5_1() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_5_1",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_5_2() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_5_2",
		&returns,
	)
	return returns
}

func H264Level_H264_LEVEL_AUTO() H264Level {
	_init_.Initialize()
	var returns H264Level
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Level",
		"H264_LEVEL_AUTO",
		&returns,
	)
	return returns
}

