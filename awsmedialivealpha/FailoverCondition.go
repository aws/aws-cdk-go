package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A condition that, when met on the active input, triggers automatic input failover to the secondary input.
//
// Create conditions with the static factory methods.
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
type FailoverCondition interface {
}

// The jsii proxy struct for FailoverCondition
type jsiiProxy_FailoverCondition struct {
	_ byte // padding
}

// Experimental.
func NewFailoverCondition_Override(f FailoverCondition) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-medialive-alpha.FailoverCondition",
		nil, // no parameters
		f,
	)
}

// Fail over when the monitored audio selector is silent for the threshold period.
// Experimental.
func FailoverCondition_AudioSilence(props *AudioSilenceFailoverProps) FailoverCondition {
	_init_.Initialize()

	if err := validateFailoverCondition_AudioSilenceParameters(props); err != nil {
		panic(err)
	}
	var returns FailoverCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FailoverCondition",
		"audioSilence",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Fail over when no input is detected for the threshold period.
// Experimental.
func FailoverCondition_InputLoss(props *InputLossFailoverProps) FailoverCondition {
	_init_.Initialize()

	if err := validateFailoverCondition_InputLossParameters(props); err != nil {
		panic(err)
	}
	var returns FailoverCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FailoverCondition",
		"inputLoss",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Fail over when the content is black for the threshold period.
// Experimental.
func FailoverCondition_VideoBlack(props *VideoBlackFailoverProps) FailoverCondition {
	_init_.Initialize()

	if err := validateFailoverCondition_VideoBlackParameters(props); err != nil {
		panic(err)
	}
	var returns FailoverCondition

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FailoverCondition",
		"videoBlack",
		[]interface{}{props},
		&returns,
	)

	return returns
}

