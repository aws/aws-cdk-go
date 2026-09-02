package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS discontinuity tags.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsDiscontinuityTags := medialive_alpha.HlsDiscontinuityTags_Of(jsii.String("value"))
//
// Experimental.
type HlsDiscontinuityTags interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsDiscontinuityTags
type jsiiProxy_HlsDiscontinuityTags struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsDiscontinuityTags) Value() *string {
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
func HlsDiscontinuityTags_Of(value *string) HlsDiscontinuityTags {
	_init_.Initialize()

	if err := validateHlsDiscontinuityTags_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsDiscontinuityTags

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsDiscontinuityTags",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsDiscontinuityTags_INSERT() HlsDiscontinuityTags {
	_init_.Initialize()
	var returns HlsDiscontinuityTags
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsDiscontinuityTags",
		"INSERT",
		&returns,
	)
	return returns
}

func HlsDiscontinuityTags_NEVER_INSERT() HlsDiscontinuityTags {
	_init_.Initialize()
	var returns HlsDiscontinuityTags
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsDiscontinuityTags",
		"NEVER_INSERT",
		&returns,
	)
	return returns
}

