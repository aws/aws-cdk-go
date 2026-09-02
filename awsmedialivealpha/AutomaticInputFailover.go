package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Automatic input failover configuration for an input attachment.
//
// When the active (primary)
// input meets any of the failover conditions, MediaLive switches to the secondary input
// without restarting the channel. This is input-source redundancy, distinct from the
// pipeline redundancy provided by `ChannelClass.STANDARD`.
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
// See: https://docs.aws.amazon.com/medialive/latest/ug/feature-automatic-input-failover.html
//
// Experimental.
type AutomaticInputFailover struct {
	// The secondary input to fail over to. Must be the same input class as the primary input.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	SecondaryInput IInput `field:"required" json:"secondaryInput" yaml:"secondaryInput"`
	// How long a recovered input must remain free of failover conditions before it is considered healthy.
	//
	// Particularly relevant with `InputPreference.PRIMARY_INPUT_PREFERRED`.
	// Default: - MediaLive service default.
	//
	// Experimental.
	ErrorClearTime awscdk.Duration `field:"optional" json:"errorClearTime" yaml:"errorClearTime"`
	// The conditions that trigger failover to the secondary input.
	// Default: - a single input-loss condition with the MediaLive service default threshold.
	//
	// Experimental.
	FailoverConditions *[]FailoverCondition `field:"optional" json:"failoverConditions" yaml:"failoverConditions"`
	// Which input to prefer once a failed input has recovered.
	// Default: - EQUAL_INPUT_PREFERENCE, applied by MediaLive.
	//
	// Experimental.
	InputPreference InputPreference `field:"optional" json:"inputPreference" yaml:"inputPreference"`
}

