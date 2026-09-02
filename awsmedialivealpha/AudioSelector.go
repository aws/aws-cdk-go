package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An audio selector that identifies which audio to extract from the input.
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
type AudioSelector interface {
	// The name of this audio selector.
	//
	// Reference it from features that monitor a specific
	// selector, such as an audio-silence {@link FailoverCondition }.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for AudioSelector
type jsiiProxy_AudioSelector struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioSelector) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewAudioSelector_Override(a AudioSelector, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		[]interface{}{name},
		a,
	)
}

// Select audio by language code.
// Experimental.
func AudioSelector_ByLanguage(name *string, languageCode *string, languageSelectionPolicy AudioLanguageSelectionPolicy) AudioSelector {
	_init_.Initialize()

	if err := validateAudioSelector_ByLanguageParameters(name, languageCode); err != nil {
		panic(err)
	}
	var returns AudioSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		"byLanguage",
		[]interface{}{name, languageCode, languageSelectionPolicy},
		&returns,
	)

	return returns
}

// Select one or more audio PIDs from the source.
// Experimental.
func AudioSelector_ByPid(name *string, pids *[]*AudioPidConfig) AudioSelector {
	_init_.Initialize()

	if err := validateAudioSelector_ByPidParameters(name, pids); err != nil {
		panic(err)
	}
	var returns AudioSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		"byPid",
		[]interface{}{name, pids},
		&returns,
	)

	return returns
}

// Select one or more audio tracks (1-based) from the source.
// Experimental.
func AudioSelector_ByTrack(name *string, tracks *[]*AudioTrackConfig, dolbyEProgramSelection DolbyEProgramSelection) AudioSelector {
	_init_.Initialize()

	if err := validateAudioSelector_ByTrackParameters(name, tracks); err != nil {
		panic(err)
	}
	var returns AudioSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		"byTrack",
		[]interface{}{name, tracks, dolbyEProgramSelection},
		&returns,
	)

	return returns
}

// Select the default audio track (no specific selector settings).
// Experimental.
func AudioSelector_Default(name *string) AudioSelector {
	_init_.Initialize()

	if err := validateAudioSelector_DefaultParameters(name); err != nil {
		panic(err)
	}
	var returns AudioSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		"default",
		[]interface{}{name},
		&returns,
	)

	return returns
}

// Select an HLS audio rendition by its `#EXT-X-MEDIA` `GROUP-ID` and `NAME`.
// Experimental.
func AudioSelector_HlsRendition(name *string, options *HlsRenditionSelectionOptions) AudioSelector {
	_init_.Initialize()

	if err := validateAudioSelector_HlsRenditionParameters(name, options); err != nil {
		panic(err)
	}
	var returns AudioSelector

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSelector",
		"hlsRendition",
		[]interface{}{name, options},
		&returns,
	)

	return returns
}

