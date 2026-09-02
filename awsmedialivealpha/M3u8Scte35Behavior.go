package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SCTE-35 passthrough behavior for an M3U8 container.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   	},
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   			},
//   			OutputName: jsii.String("video"),
//   			HlsSettings: medialive.HlsSettings_Standard(&StandardHlsSettingsProps{
//   				M3u8Settings: medialive.M3u8Settings_Of(&M3u8SettingsProps{
//   					Scte35Behavior: medialive.M3u8Scte35Behavior_PASSTHROUGH(),
//   					ProgramNum: jsii.Number(1),
//   				}),
//   			}),
//   		},
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				audio,
//   			},
//   			OutputName: jsii.String("audio"),
//   			HlsSettings: medialive.HlsSettings_AudioOnly(&AudioOnlyHlsSettingsProps{
//   				AudioGroupId: jsii.String("program"),
//   				AudioOnlyImage: medialive.FileLocation_FromBucket(bucket, jsii.String("art/cover.png")),
//   			}),
//   		},
//   	},
//   })
//
// Experimental.
type M3u8Scte35Behavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M3u8Scte35Behavior
type jsiiProxy_M3u8Scte35Behavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M3u8Scte35Behavior) Value() *string {
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
func M3u8Scte35Behavior_Of(value *string) M3u8Scte35Behavior {
	_init_.Initialize()

	if err := validateM3u8Scte35Behavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M3u8Scte35Behavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M3u8Scte35Behavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M3u8Scte35Behavior_NO_PASSTHROUGH() M3u8Scte35Behavior {
	_init_.Initialize()
	var returns M3u8Scte35Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8Scte35Behavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func M3u8Scte35Behavior_PASSTHROUGH() M3u8Scte35Behavior {
	_init_.Initialize()
	var returns M3u8Scte35Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M3u8Scte35Behavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

