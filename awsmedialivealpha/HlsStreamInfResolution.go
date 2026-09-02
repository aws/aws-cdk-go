package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS stream inf resolution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsStreamInfResolution := medialive_alpha.HlsStreamInfResolution_Of(jsii.String("value"))
//
// Experimental.
type HlsStreamInfResolution interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsStreamInfResolution
type jsiiProxy_HlsStreamInfResolution struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsStreamInfResolution) Value() *string {
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
func HlsStreamInfResolution_Of(value *string) HlsStreamInfResolution {
	_init_.Initialize()

	if err := validateHlsStreamInfResolution_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsStreamInfResolution

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsStreamInfResolution",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsStreamInfResolution_EXCLUDE() HlsStreamInfResolution {
	_init_.Initialize()
	var returns HlsStreamInfResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsStreamInfResolution",
		"EXCLUDE",
		&returns,
	)
	return returns
}

func HlsStreamInfResolution_INCLUDE() HlsStreamInfResolution {
	_init_.Initialize()
	var returns HlsStreamInfResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsStreamInfResolution",
		"INCLUDE",
		&returns,
	)
	return returns
}

