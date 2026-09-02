package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AC3 bitstream mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ac3BitstreamMode := medialive_alpha.Ac3BitstreamMode_COMMENTARY()
//
// Experimental.
type Ac3BitstreamMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Ac3BitstreamMode
type jsiiProxy_Ac3BitstreamMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Ac3BitstreamMode) Value() *string {
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
func Ac3BitstreamMode_Of(value *string) Ac3BitstreamMode {
	_init_.Initialize()

	if err := validateAc3BitstreamMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Ac3BitstreamMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Ac3BitstreamMode_COMMENTARY() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"COMMENTARY",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_COMPLETE_MAIN() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"COMPLETE_MAIN",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_DIALOGUE() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"DIALOGUE",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_EMERGENCY() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"EMERGENCY",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_HEARING_IMPAIRED() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"HEARING_IMPAIRED",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_MUSIC_AND_EFFECTS() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"MUSIC_AND_EFFECTS",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_VISUALLY_IMPAIRED() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"VISUALLY_IMPAIRED",
		&returns,
	)
	return returns
}

func Ac3BitstreamMode_VOICE_OVER() Ac3BitstreamMode {
	_init_.Initialize()
	var returns Ac3BitstreamMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3BitstreamMode",
		"VOICE_OVER",
		&returns,
	)
	return returns
}

