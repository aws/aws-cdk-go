package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 DRC line mode profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3DrcLine := medialive_alpha.Eac3DrcLine_FILM_LIGHT()
//
// Experimental.
type Eac3DrcLine interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3DrcLine
type jsiiProxy_Eac3DrcLine struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3DrcLine) Value() *string {
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
func Eac3DrcLine_Of(value *string) Eac3DrcLine {
	_init_.Initialize()

	if err := validateEac3DrcLine_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3DrcLine

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3DrcLine_FILM_LIGHT() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"FILM_LIGHT",
		&returns,
	)
	return returns
}

func Eac3DrcLine_FILM_STANDARD() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"FILM_STANDARD",
		&returns,
	)
	return returns
}

func Eac3DrcLine_MUSIC_LIGHT() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"MUSIC_LIGHT",
		&returns,
	)
	return returns
}

func Eac3DrcLine_MUSIC_STANDARD() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"MUSIC_STANDARD",
		&returns,
	)
	return returns
}

func Eac3DrcLine_NONE() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"NONE",
		&returns,
	)
	return returns
}

func Eac3DrcLine_SPEECH() Eac3DrcLine {
	_init_.Initialize()
	var returns Eac3DrcLine
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DrcLine",
		"SPEECH",
		&returns,
	)
	return returns
}

