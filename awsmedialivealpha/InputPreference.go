package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Input preference when deciding which input to make active after a previously failed input has recovered.
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
type InputPreference interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputPreference
type jsiiProxy_InputPreference struct {
	_ byte // padding
}

func (j *jsiiProxy_InputPreference) Value() *string {
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
func InputPreference_Of(value *string) InputPreference {
	_init_.Initialize()

	if err := validateInputPreference_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputPreference

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputPreference",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputPreference_EQUAL_INPUT_PREFERENCE() InputPreference {
	_init_.Initialize()
	var returns InputPreference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputPreference",
		"EQUAL_INPUT_PREFERENCE",
		&returns,
	)
	return returns
}

func InputPreference_PRIMARY_INPUT_PREFERRED() InputPreference {
	_init_.Initialize()
	var returns InputPreference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputPreference",
		"PRIMARY_INPUT_PREFERRED",
		&returns,
	)
	return returns
}

