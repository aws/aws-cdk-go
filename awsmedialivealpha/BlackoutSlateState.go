package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Blackout slate state.
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
//   		},
//   	},
//   	AvailBlanking: &AvailBlanking{
//   		State: medialive.AvailBlankingState_ENABLED(),
//   		Image: medialive.FileLocation_FromBucket(bucket, jsii.String("slates/avail.png")),
//   	},
//   	BlackoutSlate: &BlackoutSlate{
//   		State: medialive.BlackoutSlateState_ENABLED(),
//   		Image: medialive.FileLocation_*FromBucket(bucket, jsii.String("slates/blackout.png")),
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
type BlackoutSlateState interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for BlackoutSlateState
type jsiiProxy_BlackoutSlateState struct {
	_ byte // padding
}

func (j *jsiiProxy_BlackoutSlateState) Value() *string {
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
func BlackoutSlateState_Of(value *string) BlackoutSlateState {
	_init_.Initialize()

	if err := validateBlackoutSlateState_OfParameters(value); err != nil {
		panic(err)
	}
	var returns BlackoutSlateState

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.BlackoutSlateState",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func BlackoutSlateState_DISABLED() BlackoutSlateState {
	_init_.Initialize()
	var returns BlackoutSlateState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BlackoutSlateState",
		"DISABLED",
		&returns,
	)
	return returns
}

func BlackoutSlateState_ENABLED() BlackoutSlateState {
	_init_.Initialize()
	var returns BlackoutSlateState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BlackoutSlateState",
		"ENABLED",
		&returns,
	)
	return returns
}

