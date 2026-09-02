package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC coding mode.
//
// Example:
//   // AAC stereo
//   aac := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("aac_stereo"),
//   	Codec: medialive.AudioCodecSettings_Aac(&AacSettingsProps{
//   		Bitrate: awscdk.Bitrate_Kbps(jsii.Number(192)),
//   		CodingMode: medialive.AacCodingMode_CODING_MODE_2_0(),
//   	}),
//   })
//
//   // AC3 5.1
//   ac3 := medialive.EncodeConfiguration_Audio(&AudioEncodeProps{
//   	Name: jsii.String("ac3_surround"),
//   	Codec: medialive.AudioCodecSettings_Ac3(&Ac3SettingsProps{
//   		Bitrate: awscdk.Bitrate_*Kbps(jsii.Number(384)),
//   		CodingMode: medialive.Ac3CodingMode_CODING_MODE_3_2_LFE(),
//   	}),
//   })
//
// Experimental.
type AacCodingMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacCodingMode
type jsiiProxy_AacCodingMode struct {
	_ byte // padding
}

func (j *jsiiProxy_AacCodingMode) Value() *string {
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
func AacCodingMode_Of(value *string) AacCodingMode {
	_init_.Initialize()

	if err := validateAacCodingMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacCodingMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacCodingMode_AD_RECEIVER_MIX() AacCodingMode {
	_init_.Initialize()
	var returns AacCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"AD_RECEIVER_MIX",
		&returns,
	)
	return returns
}

func AacCodingMode_CODING_MODE_1_0() AacCodingMode {
	_init_.Initialize()
	var returns AacCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"CODING_MODE_1_0",
		&returns,
	)
	return returns
}

func AacCodingMode_CODING_MODE_1_1() AacCodingMode {
	_init_.Initialize()
	var returns AacCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"CODING_MODE_1_1",
		&returns,
	)
	return returns
}

func AacCodingMode_CODING_MODE_2_0() AacCodingMode {
	_init_.Initialize()
	var returns AacCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"CODING_MODE_2_0",
		&returns,
	)
	return returns
}

func AacCodingMode_CODING_MODE_5_1() AacCodingMode {
	_init_.Initialize()
	var returns AacCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacCodingMode",
		"CODING_MODE_5_1",
		&returns,
	)
	return returns
}

