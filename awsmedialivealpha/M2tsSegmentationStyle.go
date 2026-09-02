package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How segmentation markers respond to avails truncating a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsSegmentationStyle := medialive_alpha.M2tsSegmentationStyle_Of(jsii.String("value"))
//
// Experimental.
type M2tsSegmentationStyle interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsSegmentationStyle
type jsiiProxy_M2tsSegmentationStyle struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsSegmentationStyle) Value() *string {
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
func M2tsSegmentationStyle_Of(value *string) M2tsSegmentationStyle {
	_init_.Initialize()

	if err := validateM2tsSegmentationStyle_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsSegmentationStyle

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationStyle",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsSegmentationStyle_MAINTAIN_CADENCE() M2tsSegmentationStyle {
	_init_.Initialize()
	var returns M2tsSegmentationStyle
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationStyle",
		"MAINTAIN_CADENCE",
		&returns,
	)
	return returns
}

func M2tsSegmentationStyle_RESET_CADENCE() M2tsSegmentationStyle {
	_init_.Initialize()
	var returns M2tsSegmentationStyle
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationStyle",
		"RESET_CADENCE",
		&returns,
	)
	return returns
}

