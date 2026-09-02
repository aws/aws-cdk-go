package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AC3 DRC profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ac3DrcProfile := medialive_alpha.Ac3DrcProfile_Of(jsii.String("value"))
//
// Experimental.
type Ac3DrcProfile interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Ac3DrcProfile
type jsiiProxy_Ac3DrcProfile struct {
	_ byte // padding
}

func (j *jsiiProxy_Ac3DrcProfile) Value() *string {
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
func Ac3DrcProfile_Of(value *string) Ac3DrcProfile {
	_init_.Initialize()

	if err := validateAc3DrcProfile_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Ac3DrcProfile

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Ac3DrcProfile",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Ac3DrcProfile_FILM_STANDARD() Ac3DrcProfile {
	_init_.Initialize()
	var returns Ac3DrcProfile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3DrcProfile",
		"FILM_STANDARD",
		&returns,
	)
	return returns
}

func Ac3DrcProfile_NONE() Ac3DrcProfile {
	_init_.Initialize()
	var returns Ac3DrcProfile
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3DrcProfile",
		"NONE",
		&returns,
	)
	return returns
}

