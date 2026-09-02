package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 Atmos DRC line mode profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3AtmosDrcLine := medialive_alpha.Eac3AtmosDrcLine_FILM_LIGHT()
//
// Experimental.
type Eac3AtmosDrcLine interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3AtmosDrcLine
type jsiiProxy_Eac3AtmosDrcLine struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3AtmosDrcLine) Value() *string {
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
func Eac3AtmosDrcLine_Of(value *string) Eac3AtmosDrcLine {
	_init_.Initialize()

	if err := validateEac3AtmosDrcLine_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3AtmosDrcLine

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3AtmosDrcLine_FILM_LIGHT() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"FILM_LIGHT",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcLine_FILM_STANDARD() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"FILM_STANDARD",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcLine_MUSIC_LIGHT() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"MUSIC_LIGHT",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcLine_MUSIC_STANDARD() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"MUSIC_STANDARD",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcLine_NONE() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"NONE",
		&returns,
	)
	return returns
}

func Eac3AtmosDrcLine_SPEECH() Eac3AtmosDrcLine {
	_init_.Initialize()
	var returns Eac3AtmosDrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosDrcLine",
		"SPEECH",
		&returns,
	)
	return returns
}

