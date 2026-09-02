package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The audio type, as defined in ISO/IEC 13818-1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioType := medialive_alpha.AudioType_CLEAN_EFFECTS()
//
// Experimental.
type AudioType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioType
type jsiiProxy_AudioType struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioType) Value() *string {
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
func AudioType_Of(value *string) AudioType {
	_init_.Initialize()

	if err := validateAudioType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioType_CLEAN_EFFECTS() AudioType {
	_init_.Initialize()
	var returns AudioType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioType",
		"CLEAN_EFFECTS",
		&returns,
	)
	return returns
}

func AudioType_HEARING_IMPAIRED() AudioType {
	_init_.Initialize()
	var returns AudioType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioType",
		"HEARING_IMPAIRED",
		&returns,
	)
	return returns
}

func AudioType_UNDEFINED() AudioType {
	_init_.Initialize()
	var returns AudioType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioType",
		"UNDEFINED",
		&returns,
	)
	return returns
}

func AudioType_VISUAL_IMPAIRED_COMMENTARY() AudioType {
	_init_.Initialize()
	var returns AudioType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioType",
		"VISUAL_IMPAIRED_COMMENTARY",
		&returns,
	)
	return returns
}

