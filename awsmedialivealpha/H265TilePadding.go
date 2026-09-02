package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 tile padding.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265TilePadding := medialive_alpha.H265TilePadding_Of(jsii.String("value"))
//
// Experimental.
type H265TilePadding interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265TilePadding
type jsiiProxy_H265TilePadding struct {
	_ byte // padding
}

func (j *jsiiProxy_H265TilePadding) Value() *string {
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
func H265TilePadding_Of(value *string) H265TilePadding {
	_init_.Initialize()

	if err := validateH265TilePadding_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265TilePadding

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265TilePadding",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265TilePadding_NONE() H265TilePadding {
	_init_.Initialize()
	var returns H265TilePadding
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265TilePadding",
		"NONE",
		&returns,
	)
	return returns
}

func H265TilePadding_PADDED() H265TilePadding {
	_init_.Initialize()
	var returns H265TilePadding
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265TilePadding",
		"PADDED",
		&returns,
	)
	return returns
}

