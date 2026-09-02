package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Routing tier based on your maximum bitrate requirements.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//   var transitSecret Secret
//   // must hold the same value as the channel's MediaConnectRouterSettings.shared() secret
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("channel-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
//   		Channel: mediaLiveChannel,
//   		OutputName: jsii.String("router-ts"),
//   		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   		SourceTransitDecryption: &TransitEncryption{
//   			Secret: transitSecret,
//   		},
//   	}),
//   })
//
// Experimental.
type RouterInputTier interface {
	// The tier string value.
	// Experimental.
	Value() *string
	// Returns the string value.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RouterInputTier
type jsiiProxy_RouterInputTier struct {
	_ byte // padding
}

func (j *jsiiProxy_RouterInputTier) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Use a custom tier value.
// Experimental.
func RouterInputTier_Of(value *string) RouterInputTier {
	_init_.Initialize()

	if err := validateRouterInputTier_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RouterInputTier

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputTier",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RouterInputTier_INPUT_100() RouterInputTier {
	_init_.Initialize()
	var returns RouterInputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputTier",
		"INPUT_100",
		&returns,
	)
	return returns
}

func RouterInputTier_INPUT_20() RouterInputTier {
	_init_.Initialize()
	var returns RouterInputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputTier",
		"INPUT_20",
		&returns,
	)
	return returns
}

func RouterInputTier_INPUT_50() RouterInputTier {
	_init_.Initialize()
	var returns RouterInputTier
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputTier",
		"INPUT_50",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RouterInputTier) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

