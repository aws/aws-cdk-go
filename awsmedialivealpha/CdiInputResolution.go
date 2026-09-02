package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Maximum CDI input resolution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   cdiInputResolution := medialive_alpha.CdiInputResolution_FHD()
//
// Experimental.
type CdiInputResolution interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CdiInputResolution
type jsiiProxy_CdiInputResolution struct {
	_ byte // padding
}

func (j *jsiiProxy_CdiInputResolution) Value() *string {
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
func CdiInputResolution_Of(value *string) CdiInputResolution {
	_init_.Initialize()

	if err := validateCdiInputResolution_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CdiInputResolution

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CdiInputResolution",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CdiInputResolution_FHD() CdiInputResolution {
	_init_.Initialize()
	var returns CdiInputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CdiInputResolution",
		"FHD",
		&returns,
	)
	return returns
}

func CdiInputResolution_HD() CdiInputResolution {
	_init_.Initialize()
	var returns CdiInputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CdiInputResolution",
		"HD",
		&returns,
	)
	return returns
}

func CdiInputResolution_SD() CdiInputResolution {
	_init_.Initialize()
	var returns CdiInputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CdiInputResolution",
		"SD",
		&returns,
	)
	return returns
}

func CdiInputResolution_UHD() CdiInputResolution {
	_init_.Initialize()
	var returns CdiInputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CdiInputResolution",
		"UHD",
		&returns,
	)
	return returns
}

