package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS client cache control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsClientCache := medialive_alpha.HlsClientCache_Of(jsii.String("value"))
//
// Experimental.
type HlsClientCache interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsClientCache
type jsiiProxy_HlsClientCache struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsClientCache) Value() *string {
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
func HlsClientCache_Of(value *string) HlsClientCache {
	_init_.Initialize()

	if err := validateHlsClientCache_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsClientCache

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsClientCache",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsClientCache_DISABLED() HlsClientCache {
	_init_.Initialize()
	var returns HlsClientCache
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsClientCache",
		"DISABLED",
		&returns,
	)
	return returns
}

func HlsClientCache_ENABLED() HlsClientCache {
	_init_.Initialize()
	var returns HlsClientCache
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsClientCache",
		"ENABLED",
		&returns,
	)
	return returns
}

