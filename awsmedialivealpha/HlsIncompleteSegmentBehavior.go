package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS incomplete segment behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsIncompleteSegmentBehavior := medialive_alpha.HlsIncompleteSegmentBehavior_Of(jsii.String("value"))
//
// Experimental.
type HlsIncompleteSegmentBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsIncompleteSegmentBehavior
type jsiiProxy_HlsIncompleteSegmentBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsIncompleteSegmentBehavior) Value() *string {
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
func HlsIncompleteSegmentBehavior_Of(value *string) HlsIncompleteSegmentBehavior {
	_init_.Initialize()

	if err := validateHlsIncompleteSegmentBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsIncompleteSegmentBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsIncompleteSegmentBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsIncompleteSegmentBehavior_AUTO() HlsIncompleteSegmentBehavior {
	_init_.Initialize()
	var returns HlsIncompleteSegmentBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIncompleteSegmentBehavior",
		"AUTO",
		&returns,
	)
	return returns
}

func HlsIncompleteSegmentBehavior_SUPPRESS() HlsIncompleteSegmentBehavior {
	_init_.Initialize()
	var returns HlsIncompleteSegmentBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIncompleteSegmentBehavior",
		"SUPPRESS",
		&returns,
	)
	return returns
}

