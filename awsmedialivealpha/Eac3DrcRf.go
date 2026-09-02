package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 DRC RF mode profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3DrcRf := medialive_alpha.Eac3DrcRf_FILM_LIGHT()
//
// Experimental.
type Eac3DrcRf interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3DrcRf
type jsiiProxy_Eac3DrcRf struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3DrcRf) Value() *string {
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
func Eac3DrcRf_Of(value *string) Eac3DrcRf {
	_init_.Initialize()

	if err := validateEac3DrcRf_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3DrcRf

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3DrcRf_FILM_LIGHT() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"FILM_LIGHT",
		&returns,
	)
	return returns
}

func Eac3DrcRf_FILM_STANDARD() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"FILM_STANDARD",
		&returns,
	)
	return returns
}

func Eac3DrcRf_MUSIC_LIGHT() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"MUSIC_LIGHT",
		&returns,
	)
	return returns
}

func Eac3DrcRf_MUSIC_STANDARD() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"MUSIC_STANDARD",
		&returns,
	)
	return returns
}

func Eac3DrcRf_NONE() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"NONE",
		&returns,
	)
	return returns
}

func Eac3DrcRf_SPEECH() Eac3DrcRf {
	_init_.Initialize()
	var returns Eac3DrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcRf",
		"SPEECH",
		&returns,
	)
	return returns
}

