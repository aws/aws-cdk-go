package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Policy for how MediaLive identifies the audio stream when selecting by language, on a transport-stream PMT update.
//
// Example:
//   var stack Stack
//   var input IInput
//   var bucket IBucket
//   var video EncodeConfiguration
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: *Input,
//   			AudioSelectors: []AudioSelector{
//   				medialive.AudioSelector_ByLanguage(jsii.String("eng"), jsii.String("eng"), medialive.AudioLanguageSelectionPolicy_STRICT()),
//   			},
//   			CaptionSelectors: []CaptionSelector{
//   				medialive.CaptionSelector_Embedded(jsii.String("embedded")),
//   			},
//   			VideoSelector: &VideoSelectorSettings{
//   				ColorSpace: medialive.VideoColorSpace_HDR10(),
//   				ColorSpaceUsage: medialive.VideoColorSpaceUsage_FORCE(),
//   				SelectBy: medialive.VideoSelection_ByProgramId(jsii.Number(1)),
//   			},
//   		},
//   	},
//   	OutputGroups: []OutputGroupConfiguration{
//   		medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   			Name: jsii.String("hls"),
//   			Destinations: []OutputDestination{
//   				medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   			},
//   			Outputs: []HlsOutputDefinition{
//   				&HlsOutputDefinition{
//   					Encodes: []EncodeConfiguration{
//   						video,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type AudioLanguageSelectionPolicy interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioLanguageSelectionPolicy
type jsiiProxy_AudioLanguageSelectionPolicy struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioLanguageSelectionPolicy) Value() *string {
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
func AudioLanguageSelectionPolicy_Of(value *string) AudioLanguageSelectionPolicy {
	_init_.Initialize()

	if err := validateAudioLanguageSelectionPolicy_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioLanguageSelectionPolicy

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageSelectionPolicy",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioLanguageSelectionPolicy_LOOSE() AudioLanguageSelectionPolicy {
	_init_.Initialize()
	var returns AudioLanguageSelectionPolicy
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageSelectionPolicy",
		"LOOSE",
		&returns,
	)
	return returns
}

func AudioLanguageSelectionPolicy_STRICT() AudioLanguageSelectionPolicy {
	_init_.Initialize()
	var returns AudioLanguageSelectionPolicy
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageSelectionPolicy",
		"STRICT",
		&returns,
	)
	return returns
}

