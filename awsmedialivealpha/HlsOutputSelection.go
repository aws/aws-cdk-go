package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS output selection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsOutputSelection := medialive_alpha.HlsOutputSelection_MANIFESTS_AND_SEGMENTS()
//
// Experimental.
type HlsOutputSelection interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsOutputSelection
type jsiiProxy_HlsOutputSelection struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsOutputSelection) Value() *string {
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
func HlsOutputSelection_Of(value *string) HlsOutputSelection {
	_init_.Initialize()

	if err := validateHlsOutputSelection_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsOutputSelection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsOutputSelection",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsOutputSelection_MANIFESTS_AND_SEGMENTS() HlsOutputSelection {
	_init_.Initialize()
	var returns HlsOutputSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsOutputSelection",
		"MANIFESTS_AND_SEGMENTS",
		&returns,
	)
	return returns
}

func HlsOutputSelection_SEGMENTS_ONLY() HlsOutputSelection {
	_init_.Initialize()
	var returns HlsOutputSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsOutputSelection",
		"SEGMENTS_ONLY",
		&returns,
	)
	return returns
}

func HlsOutputSelection_VARIANT_MANIFESTS_AND_SEGMENTS() HlsOutputSelection {
	_init_.Initialize()
	var returns HlsOutputSelection
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsOutputSelection",
		"VARIANT_MANIFESTS_AND_SEGMENTS",
		&returns,
	)
	return returns
}

