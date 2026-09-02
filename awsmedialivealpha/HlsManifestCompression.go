package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS manifest compression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsManifestCompression := medialive_alpha.HlsManifestCompression_Of(jsii.String("value"))
//
// Experimental.
type HlsManifestCompression interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsManifestCompression
type jsiiProxy_HlsManifestCompression struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsManifestCompression) Value() *string {
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
func HlsManifestCompression_Of(value *string) HlsManifestCompression {
	_init_.Initialize()

	if err := validateHlsManifestCompression_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsManifestCompression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsManifestCompression",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsManifestCompression_GZIP() HlsManifestCompression {
	_init_.Initialize()
	var returns HlsManifestCompression
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsManifestCompression",
		"GZIP",
		&returns,
	)
	return returns
}

func HlsManifestCompression_NONE() HlsManifestCompression {
	_init_.Initialize()
	var returns HlsManifestCompression
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsManifestCompression",
		"NONE",
		&returns,
	)
	return returns
}

