package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 Atmos DRC RF mode profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3AtmosDrcRf := medialive_alpha.Eac3AtmosDrcRf_FILM_LIGHT()
//
// Experimental.
type Eac3AtmosDrcRf interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3AtmosDrcRf
type jsiiProxy_Eac3AtmosDrcRf struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3AtmosDrcRf) Value() *string {
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
func Eac3AtmosDrcRf_Of(value *string) Eac3AtmosDrcRf {
	_init_.Initialize()

	if err := validateEac3AtmosDrcRf_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3AtmosDrcRf

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3AtmosDrcRf_FILM_LIGHT() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"FILM_LIGHT",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcRf_FILM_STANDARD() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"FILM_STANDARD",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcRf_MUSIC_LIGHT() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"MUSIC_LIGHT",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcRf_MUSIC_STANDARD() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"MUSIC_STANDARD",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcRf_NONE() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"NONE",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcRf_SPEECH() Eac3AtmosDrcRf {
	_init_.Initialize()
	var returns Eac3AtmosDrcRf
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcRf",
		"SPEECH",
		&returns,
	)
	return returns
}

