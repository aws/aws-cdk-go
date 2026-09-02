package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Router Output configurations.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("SrtOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("srt-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	// tier defaults to RouterOutputTier.OUTPUT_20 (lowest cost)
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   		Protocol: awsmediaconnectalpha.RouterOutputProtocol_SrtListener(&SrtListenerOutputProtocolProps{
//   			Port: jsii.Number(9001),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   		}),
//   		NetworkInterface: networkInterface,
//   	}),
//   })
//
// Experimental.
type RouterOutputConfiguration interface {
}

// The jsii proxy struct for RouterOutputConfiguration
type jsiiProxy_RouterOutputConfiguration struct {
	_ byte // padding
}

// Experimental.
func NewRouterOutputConfiguration_Override(r RouterOutputConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		nil, // no parameters
		r,
	)
}

// Create a MediaConnect Flow Router Output configuration with a specific flow connection.
//
// Use this when the target flow already exists and you want to connect immediately.
//
// Returns: RouterOutputConfiguration instance for MediaConnect Flow setup with flow connection.
//
// Example:
//   var flow Flow
//
//
//   awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlow(&MediaConnectFlowConnectionProps{
//   	Flow: Flow,
//   })
//
// Experimental.
func RouterOutputConfiguration_MediaConnectFlow(props *MediaConnectFlowConnectionProps) RouterOutputConfiguration {
	_init_.Initialize()

	if err := validateRouterOutputConfiguration_MediaConnectFlowParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		"mediaConnectFlow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaConnect Flow Router Output configuration without a specific flow connection.
//
// Use this when you want to set up the router output before the target flow exists.
//
// Returns: RouterOutputConfiguration instance for MediaConnect Flow setup without flow connection.
//
// Example:
//   awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlowWithoutConnection(&MediaConnectFlowNoConnectionProps{
//   	AvailabilityZone: jsii.String("us-east-1a"),
//   })
//
// Experimental.
func RouterOutputConfiguration_MediaConnectFlowWithoutConnection(props *MediaConnectFlowNoConnectionProps) RouterOutputConfiguration {
	_init_.Initialize()

	if err := validateRouterOutputConfiguration_MediaConnectFlowWithoutConnectionParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		"mediaConnectFlowWithoutConnection",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaLive Router Output configuration with a specific input connection.
//
// Use this when the MediaLive input already exists and you want to connect immediately.
// Experimental.
func RouterOutputConfiguration_MediaLiveInput(props *MediaLiveInputConnectionProps) RouterOutputConfiguration {
	_init_.Initialize()

	if err := validateRouterOutputConfiguration_MediaLiveInputParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		"mediaLiveInput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a MediaLive Router Output configuration without a specific input connection.
//
// Use this when you want to set up the router output before the MediaLive input exists.
// Experimental.
func RouterOutputConfiguration_MediaLiveInputWithoutConnection(props *MediaLiveNoInputConnectionProps) RouterOutputConfiguration {
	_init_.Initialize()

	if err := validateRouterOutputConfiguration_MediaLiveInputWithoutConnectionParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		"mediaLiveInputWithoutConnection",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a standard Router Output configuration with a single protocol.
//
// Returns: RouterOutputConfiguration instance for standard setup.
//
// Example:
//   var networkInterface RouterNetworkInterface
//
//
//   awsmediaconnectalpha.RouterOutputConfiguration_Standard(&StandardOutputConfigurationProps{
//   	NetworkInterface: NetworkInterface,
//   	Protocol: awsmediaconnectalpha.RouterOutputProtocol_Rtp(&RtpOutputProtocolProps{
//   		DestinationAddress: jsii.String("10.0.0.1"),
//   		Port: jsii.Number(5000),
//   	}),
//   })
//
// Experimental.
func RouterOutputConfiguration_Standard(props *StandardOutputConfigurationProps) RouterOutputConfiguration {
	_init_.Initialize()

	if err := validateRouterOutputConfiguration_StandardParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputConfiguration",
		"standard",
		[]interface{}{props},
		&returns,
	)

	return returns
}

