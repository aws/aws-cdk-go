package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS manifest duration format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsManifestDurationFormat := medialive_alpha.HlsManifestDurationFormat_Of(jsii.String("value"))
//
// Experimental.
type HlsManifestDurationFormat interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsManifestDurationFormat
type jsiiProxy_HlsManifestDurationFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsManifestDurationFormat) Value() *string {
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
func HlsManifestDurationFormat_Of(value *string) HlsManifestDurationFormat {
	_init_.Initialize()

	if err := validateHlsManifestDurationFormat_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsManifestDurationFormat

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsManifestDurationFormat",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsManifestDurationFormat_FLOATING_POINT() HlsManifestDurationFormat {
	_init_.Initialize()
	var returns HlsManifestDurationFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsManifestDurationFormat",
		"FLOATING_POINT",
		&returns,
	)
	return returns
}

func HlsManifestDurationFormat_INTEGER() HlsManifestDurationFormat {
	_init_.Initialize()
	var returns HlsManifestDurationFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsManifestDurationFormat",
		"INTEGER",
		&returns,
	)
	return returns
}

