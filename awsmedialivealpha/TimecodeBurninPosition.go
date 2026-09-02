package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Position for timecode burn-in overlay.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   timecodeBurninPosition := medialive_alpha.TimecodeBurninPosition_BOTTOM_CENTER()
//
// Experimental.
type TimecodeBurninPosition interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for TimecodeBurninPosition
type jsiiProxy_TimecodeBurninPosition struct {
	_ byte // padding
}

func (j *jsiiProxy_TimecodeBurninPosition) Value() *string {
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
func TimecodeBurninPosition_Of(value *string) TimecodeBurninPosition {
	_init_.Initialize()

	if err := validateTimecodeBurninPosition_OfParameters(value); err != nil {
		panic(err)
	}
	var returns TimecodeBurninPosition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TimecodeBurninPosition_BOTTOM_CENTER() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"BOTTOM_CENTER",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_BOTTOM_LEFT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"BOTTOM_LEFT",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_BOTTOM_RIGHT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"BOTTOM_RIGHT",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_MIDDLE_CENTER() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"MIDDLE_CENTER",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_MIDDLE_LEFT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"MIDDLE_LEFT",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_MIDDLE_RIGHT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"MIDDLE_RIGHT",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_TOP_CENTER() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"TOP_CENTER",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_TOP_LEFT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"TOP_LEFT",
		&returns,
	)
	return returns
}

func TimecodeBurninPosition_TOP_RIGHT() TimecodeBurninPosition {
	_init_.Initialize()
	var returns TimecodeBurninPosition
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.TimecodeBurninPosition",
		"TOP_RIGHT",
		&returns,
	)
	return returns
}

