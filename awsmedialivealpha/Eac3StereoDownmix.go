package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 stereo downmix preference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3StereoDownmix := medialive_alpha.Eac3StereoDownmix_DPL2()
//
// Experimental.
type Eac3StereoDownmix interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3StereoDownmix
type jsiiProxy_Eac3StereoDownmix struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3StereoDownmix) Value() *string {
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
func Eac3StereoDownmix_Of(value *string) Eac3StereoDownmix {
	_init_.Initialize()

	if err := validateEac3StereoDownmix_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3StereoDownmix

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3StereoDownmix",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3StereoDownmix_DPL2() Eac3StereoDownmix {
	_init_.Initialize()
	var returns Eac3StereoDownmix
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3StereoDownmix",
		"DPL2",
		&returns,
	)
	return returns
}

func Eac3StereoDownmix_LO_RO() Eac3StereoDownmix {
	_init_.Initialize()
	var returns Eac3StereoDownmix
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3StereoDownmix",
		"LO_RO",
		&returns,
	)
	return returns
}

func Eac3StereoDownmix_LT_RT() Eac3StereoDownmix {
	_init_.Initialize()
	var returns Eac3StereoDownmix
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3StereoDownmix",
		"LT_RT",
		&returns,
	)
	return returns
}

func Eac3StereoDownmix_NOT_INDICATED() Eac3StereoDownmix {
	_init_.Initialize()
	var returns Eac3StereoDownmix
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3StereoDownmix",
		"NOT_INDICATED",
		&returns,
	)
	return returns
}

