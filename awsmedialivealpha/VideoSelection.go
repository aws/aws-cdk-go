package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Selects the specific video to extract from the input — by PID or by program.
//
// Create with
// the static factory methods; exactly one selection applies, enforced by the type.
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
type VideoSelection interface {
}

// The jsii proxy struct for VideoSelection
type jsiiProxy_VideoSelection struct {
	_ byte // padding
}

// Experimental.
func NewVideoSelection_Override(v VideoSelection) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.VideoSelection",
		nil, // no parameters
		v,
	)
}

// Extract the video with this PID.
// Experimental.
func VideoSelection_ByPid(pid *float64) VideoSelection {
	_init_.Initialize()

	if err := validateVideoSelection_ByPidParameters(pid); err != nil {
		panic(err)
	}
	var returns VideoSelection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoSelection",
		"byPid",
		[]interface{}{pid},
		&returns,
	)

	return returns
}

// Extract the video from this program within a multi-program transport stream.
//
// If the
// program doesn't exist, MediaLive selects the first program in the stream.
// Experimental.
func VideoSelection_ByProgramId(programId *float64) VideoSelection {
	_init_.Initialize()

	if err := validateVideoSelection_ByProgramIdParameters(programId); err != nil {
		panic(err)
	}
	var returns VideoSelection

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.VideoSelection",
		"byProgramId",
		[]interface{}{programId},
		&returns,
	)

	return returns
}

