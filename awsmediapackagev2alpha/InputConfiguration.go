package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Input configuration for a MediaPackage V2 Channel.
//
// Use the static factory methods to create instances:
// - InputConfiguration.hls() for HLS input
// - InputConfiguration.cmaf() for CMAF input with optional CMAF-specific features
//
// Example:
//   var stack Stack
//
//   group := awsmediapackagev2alpha.NewChannelGroup(stack, jsii.String("MyChannelGroup"), &ChannelGroupProps{
//   	ChannelGroupName: jsii.String("my-test-channel-group"),
//   })
//
//   channel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("MyChannel"), &ChannelProps{
//   	ChannelGroup: group,
//   	ChannelName: jsii.String("my-testchannel"),
//   	Input: awsmediapackagev2alpha.InputConfiguration_Cmaf(),
//   })
//
//   endpoint := awsmediapackagev2alpha.NewOriginEndpoint(stack, jsii.String("MyOriginEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	OriginEndpointName: jsii.String("my-test-endpoint"),
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type InputConfiguration interface {
	// Input switch configuration (CMAF only).
	// Experimental.
	InputSwitchConfiguration() *InputSwitchConfiguration
	// The input type (HLS or CMAF).
	// Experimental.
	InputType() InputType
	// Output headers configuration (CMAF only).
	// Experimental.
	OutputHeaders() *[]HeadersCMSD
}

// The jsii proxy struct for InputConfiguration
type jsiiProxy_InputConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_InputConfiguration) InputSwitchConfiguration() *InputSwitchConfiguration {
	var returns *InputSwitchConfiguration
	_jsii_.Get(
		j,
		"inputSwitchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputConfiguration) InputType() InputType {
	var returns InputType
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InputConfiguration) OutputHeaders() *[]HeadersCMSD {
	var returns *[]HeadersCMSD
	_jsii_.Get(
		j,
		"outputHeaders",
		&returns,
	)
	return returns
}


// Create a CMAF input configuration.
// Experimental.
func InputConfiguration_Cmaf(props *CmafInputProps) InputConfiguration {
	_init_.Initialize()

	if err := validateInputConfiguration_CmafParameters(props); err != nil {
		panic(err)
	}
	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.InputConfiguration",
		"cmaf",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an HLS input configuration.
// Experimental.
func InputConfiguration_Hls() InputConfiguration {
	_init_.Initialize()

	var returns InputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.InputConfiguration",
		"hls",
		nil, // no parameters
		&returns,
	)

	return returns
}

