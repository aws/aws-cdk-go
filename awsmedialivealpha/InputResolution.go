package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The resolution for the input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputResolution := medialive_alpha.InputResolution_HD()
//
// Experimental.
type InputResolution interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputResolution
type jsiiProxy_InputResolution struct {
	_ byte // padding
}

func (j *jsiiProxy_InputResolution) Value() *string {
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
func InputResolution_Of(value *string) InputResolution {
	_init_.Initialize()

	if err := validateInputResolution_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputResolution

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputResolution",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputResolution_HD() InputResolution {
	_init_.Initialize()
	var returns InputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputResolution",
		"HD",
		&returns,
	)
	return returns
}

func InputResolution_SD() InputResolution {
	_init_.Initialize()
	var returns InputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputResolution",
		"SD",
		&returns,
	)
	return returns
}

func InputResolution_UHD() InputResolution {
	_init_.Initialize()
	var returns InputResolution
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputResolution",
		"UHD",
		&returns,
	)
	return returns
}

