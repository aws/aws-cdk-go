package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS IV in manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsIvInManifest := medialive_alpha.HlsIvInManifest_Of(jsii.String("value"))
//
// Experimental.
type HlsIvInManifest interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsIvInManifest
type jsiiProxy_HlsIvInManifest struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsIvInManifest) Value() *string {
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
func HlsIvInManifest_Of(value *string) HlsIvInManifest {
	_init_.Initialize()

	if err := validateHlsIvInManifest_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsIvInManifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsIvInManifest",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsIvInManifest_EXCLUDE() HlsIvInManifest {
	_init_.Initialize()
	var returns HlsIvInManifest
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIvInManifest",
		"EXCLUDE",
		&returns,
	)
	return returns
}

func HlsIvInManifest_INCLUDE() HlsIvInManifest {
	_init_.Initialize()
	var returns HlsIvInManifest
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIvInManifest",
		"INCLUDE",
		&returns,
	)
	return returns
}

