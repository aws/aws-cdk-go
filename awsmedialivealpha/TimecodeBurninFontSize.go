package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Font size for timecode burn-in.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   timecodeBurninFontSize := medialive_alpha.TimecodeBurninFontSize_EXTRA_SMALL_10()
//
// Experimental.
type TimecodeBurninFontSize interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimecodeBurninFontSize
type jsiiProxy_TimecodeBurninFontSize struct {
	_ byte // padding
}

func (j *jsiiProxy_TimecodeBurninFontSize) Value() *string {
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
func TimecodeBurninFontSize_Of(value *string) TimecodeBurninFontSize {
	_init_.Initialize()

	if err := validateTimecodeBurninFontSize_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimecodeBurninFontSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninFontSize",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimecodeBurninFontSize_EXTRA_SMALL_10() TimecodeBurninFontSize {
	_init_.Initialize()
	var returns TimecodeBurninFontSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninFontSize",
		"EXTRA_SMALL_10",
		&returns,
	)
	return returns
}

func TimecodeBurninFontSize_LARGE_48() TimecodeBurninFontSize {
	_init_.Initialize()
	var returns TimecodeBurninFontSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninFontSize",
		"LARGE_48",
		&returns,
	)
	return returns
}

func TimecodeBurninFontSize_MEDIUM_32() TimecodeBurninFontSize {
	_init_.Initialize()
	var returns TimecodeBurninFontSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninFontSize",
		"MEDIUM_32",
		&returns,
	)
	return returns
}

func TimecodeBurninFontSize_SMALL_16() TimecodeBurninFontSize {
	_init_.Initialize()
	var returns TimecodeBurninFontSize
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninFontSize",
		"SMALL_16",
		&returns,
	)
	return returns
}

