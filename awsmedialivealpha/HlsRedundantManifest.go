package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS redundant manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsRedundantManifest := medialive_alpha.HlsRedundantManifest_Of(jsii.String("value"))
//
// Experimental.
type HlsRedundantManifest interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsRedundantManifest
type jsiiProxy_HlsRedundantManifest struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsRedundantManifest) Value() *string {
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
func HlsRedundantManifest_Of(value *string) HlsRedundantManifest {
	_init_.Initialize()

	if err := validateHlsRedundantManifest_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsRedundantManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsRedundantManifest",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsRedundantManifest_DISABLED() HlsRedundantManifest {
	_init_.Initialize()
	var returns HlsRedundantManifest
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsRedundantManifest",
		"DISABLED",
		&returns,
	)
	return returns
}

func HlsRedundantManifest_ENABLED() HlsRedundantManifest {
	_init_.Initialize()
	var returns HlsRedundantManifest
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsRedundantManifest",
		"ENABLED",
		&returns,
	)
	return returns
}

