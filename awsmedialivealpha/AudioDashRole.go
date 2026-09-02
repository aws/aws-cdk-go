package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A DASH role to assign to an audio output (used when the output carries DVB DASH accessibility signaling).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioDashRole := medialive_alpha.AudioDashRole_ALTERNATE()
//
// Experimental.
type AudioDashRole interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioDashRole
type jsiiProxy_AudioDashRole struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioDashRole) Value() *string {
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
func AudioDashRole_Of(value *string) AudioDashRole {
	_init_.Initialize()

	if err := validateAudioDashRole_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioDashRole

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioDashRole_ALTERNATE() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"ALTERNATE",
		&returns,
	)
	return returns
}

func AudioDashRole_COMMENTARY() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"COMMENTARY",
		&returns,
	)
	return returns
}

func AudioDashRole_DESCRIPTION() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"DESCRIPTION",
		&returns,
	)
	return returns
}

func AudioDashRole_DUB() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"DUB",
		&returns,
	)
	return returns
}

func AudioDashRole_EMERGENCY() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"EMERGENCY",
		&returns,
	)
	return returns
}

func AudioDashRole_ENHANCED_AUDIO_INTELLIGIBILITY() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"ENHANCED_AUDIO_INTELLIGIBILITY",
		&returns,
	)
	return returns
}

func AudioDashRole_KARAOKE() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"KARAOKE",
		&returns,
	)
	return returns
}

func AudioDashRole_MAIN() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"MAIN",
		&returns,
	)
	return returns
}

func AudioDashRole_SUPPLEMENTARY() AudioDashRole {
	_init_.Initialize()
	var returns AudioDashRole
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioDashRole",
		"SUPPLEMENTARY",
		&returns,
	)
	return returns
}

