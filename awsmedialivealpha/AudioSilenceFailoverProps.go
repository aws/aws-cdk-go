package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for an audio-silence failover condition.
//
// Example:
//   var stack Stack
//   var primaryInput IInput
//   var secondaryInput IInput
//   var audioSelector AudioSelector
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//   var bucket IBucket
//
//
//   medialive.NewChannel(stack, jsii.String("Channel"), &ChannelProps{
//   	Inputs: []InputAttachment{
//   		&InputAttachment{
//   			Input: primaryInput,
//   			AutomaticInputFailover: &AutomaticInputFailover{
//   				SecondaryInput: *SecondaryInput,
//   				InputPreference: medialive.InputPreference_PRIMARY_INPUT_PREFERRED(),
//   				ErrorClearTime: awscdk.Duration_Seconds(jsii.Number(3)),
//   				FailoverConditions: []FailoverCondition{
//   					medialive.FailoverCondition_InputLoss(&InputLossFailoverProps{
//   						Threshold: awscdk.Duration_Millis(jsii.Number(1500)),
//   					}),
//   					medialive.FailoverCondition_AudioSilence(&AudioSilenceFailoverProps{
//   						AudioSelector: *AudioSelector,
//   						Threshold: awscdk.Duration_*Seconds(jsii.Number(2)),
//   					}),
//   					medialive.FailoverCondition_VideoBlack(&VideoBlackFailoverProps{
//   						BlackDetectThreshold: jsii.Number(0.1),
//   						Threshold: awscdk.Duration_*Seconds(jsii.Number(1)),
//   					}),
//   				},
//   			},
//   		},
//   		&InputAttachment{
//   			// The secondary input must also be attached to the channel as its own input.
//   			Input: secondaryInput,
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
//   						audio,
//   					},
//   					OutputName: jsii.String("hls_out"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type AudioSilenceFailoverProps struct {
	// The audio selector (defined on the input attachment) that MediaLive monitors for silence.
	//
	// Pass the same `AudioSelector` instance you put in the attachment's
	// `audioSelectors` — select your most important rendition.
	// Experimental.
	AudioSelector AudioSelector `field:"required" json:"audioSelector" yaml:"audioSelector"`
	// How long the audio must be silent (audio loss, or quieter than -50 dBFS) before failover occurs.
	// Default: - MediaLive service default.
	//
	// Experimental.
	Threshold awscdk.Duration `field:"optional" json:"threshold" yaml:"threshold"`
}

