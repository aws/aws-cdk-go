package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// GOP size (keyframe interval). Use the static factory methods to specify in frames or seconds.
//
// The value must be greater than zero. When expressed in frames it must be a whole number;
// when expressed in seconds it may be fractional.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   gopSize := medialive_alpha.GopSize_Frames(jsii.Number(123))
//
// Experimental.
type GopSize interface {
}

// The jsii proxy struct for GopSize
type jsiiProxy_GopSize struct {
	_ byte // padding
}

// GOP size in frames.
//
// Must be a whole number.
// Experimental.
func GopSize_Frames(value *float64) GopSize {
	_init_.Initialize()

	if err := validateGopSize_FramesParameters(value); err != nil {
		panic(err)
	}
	var returns GopSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.GopSize",
		"frames",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// GOP size in seconds.
//
// May be fractional (e.g. `1.5`).
// Experimental.
func GopSize_Seconds(value *float64) GopSize {
	_init_.Initialize()

	if err := validateGopSize_SecondsParameters(value); err != nil {
		panic(err)
	}
	var returns GopSize

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.GopSize",
		"seconds",
		[]interface{}{value},
		&returns,
	)

	return returns
}

