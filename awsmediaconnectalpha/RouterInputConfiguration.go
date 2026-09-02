package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Router Input configurations.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//
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
//   	}),
//   })
//
// Experimental.
type RouterInputConfiguration interface {
}

// The jsii proxy struct for RouterInputConfiguration
type jsiiProxy_RouterInputConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewRouterInputConfiguration_Override(r RouterInputConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		nil, // no parameters
		r,
	)
}

// Create a failover Router Input configuration with two matching protocols.
//
// Returns: RouterInputConfiguration instance for failover setup.
// Experimental.
func RouterInputConfiguration_Failover(props *FailoverConfigurationProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_FailoverParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"failover",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect Flow Router Input configuration.
//
// Use this when the source flow already exists and you want to connect immediately.
//
// Returns: RouterInputConfiguration instance for MediaConnect Flow setup.
// Experimental.
func RouterInputConfiguration_MediaConnectFlow(props *MediaConnectFlowConfigurationProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_MediaConnectFlowParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"mediaConnectFlow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect Flow Router Input configuration without connecting to a specific flow Use this when you want to prepare the router input for a flow connection later.
//
// Returns: RouterInputConfiguration instance for MediaConnect Flow setup without connection.
// Experimental.
func RouterInputConfiguration_MediaConnectFlowWithoutConnection(props *MediaConnectFlowConfigurationWithoutConnectionProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_MediaConnectFlowWithoutConnectionParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"mediaConnectFlowWithoutConnection",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaLive Channel Router Input configuration.
//
// Use this when the source MediaLive channel already exists and you want to
// ingest from one of its outputs immediately.
//
// Returns: RouterInputConfiguration instance for MediaLive channel setup.
// Experimental.
func RouterInputConfiguration_MediaLiveChannel(props *MediaLiveChannelConfigurationProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_MediaLiveChannelParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"mediaLiveChannel",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaLive Channel Router Input configuration without a specific channel connection.
//
// Use this when you want to set up the router input before the target MediaLive channel exists.
//
// Returns: RouterInputConfiguration instance for MediaLive channel setup without connection.
// Experimental.
func RouterInputConfiguration_MediaLiveChannelWithoutConnection(props *MediaLiveChannelConfigurationWithoutConnectionProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_MediaLiveChannelWithoutConnectionParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"mediaLiveChannelWithoutConnection",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a merge Router Input configuration with two matching non-SRT protocols.
//
// Returns: RouterInputConfiguration instance for merge setup.
// Experimental.
func RouterInputConfiguration_Merge(props *MergeConfigurationProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_MergeParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"merge",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a standard Router Input configuration with a single protocol.
//
// Returns: RouterInputConfiguration instance for standard setup.
// Experimental.
func RouterInputConfiguration_Standard(props *StandardConfigurationProps) RouterInputConfiguration {
	_init_.Initialize()

	if err := validateRouterInputConfiguration_StandardParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputConfiguration",
		"standard",
		[]interface{}{props},
		&returns,
	)

	return returns
}

