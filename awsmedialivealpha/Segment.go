package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The length of a media segment for an output group.
//
// Express the length in whole seconds with {@link Segment.seconds} or in milliseconds
// with {@link Segment.milliseconds}. Some output groups (e.g. HLS) only support
// whole-second segments and will reject sub-second millisecond values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   segment := medialive_alpha.Segment_Milliseconds(jsii.Number(123))
//
// Experimental.
type Segment interface {
}

// The jsii proxy struct for Segment
type jsiiProxy_Segment struct {
	_ byte // padding
}

// A segment length in milliseconds.
// Experimental.
func Segment_Milliseconds(value *float64) Segment {
	_init_.Initialize()

	if err := validateSegment_MillisecondsParameters(value); err != nil {
		panic(err)
	}
	var returns Segment

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Segment",
		"milliseconds",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// A segment length in whole seconds.
// Experimental.
func Segment_Seconds(value *float64) Segment {
	_init_.Initialize()

	if err := validateSegment_SecondsParameters(value); err != nil {
		panic(err)
	}
	var returns Segment

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Segment",
		"seconds",
		[]interface{}{value},
		&returns,
	)

	return returns
}

