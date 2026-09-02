package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Router Input protocol configurations.
//
// Example:
//   var stack Stack
//   var mediaLiveInput IInput
//
//
//   // 1. A public network interface for the SRT input
//   networkInterface := awsmediaconnectalpha.NewRouterNetworkInterface(stack, jsii.String("NetworkInterface"), &RouterNetworkInterfaceProps{
//   	RouterNetworkInterfaceName: jsii.String("camera-network"),
//   	Configuration: awsmediaconnectalpha.RouterNetworkConfiguration_PublicNetwork(&PublicNetworkConfigurationProps{
//   		Cidr: []*string{
//   			jsii.String("203.0.113.0/24"),
//   		},
//   	}),
//   })
//
//   // 2. A router input receiving SRT from an upstream encoder
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("Input"), &RouterInputProps{
//   	RouterInputName: jsii.String("camera-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_20(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Standard(&StandardConfigurationProps{
//   		NetworkInterface: *NetworkInterface,
//   		Protocol: awsmediaconnectalpha.RouterInputProtocol_SrtListener(&SrtListenerProtocolProps{
//   			Port: jsii.Number(9000),
//   			MinimumLatency: awscdk.Duration_Millis(jsii.Number(200)),
//   		}),
//   	}),
//   })
//
//   // 3. A router output delivering to MediaLive
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("Output"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("medialive-output"),
//   	MaximumBitrate: awscdk.Bitrate_*Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_20(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInput(&MediaLiveInputConnectionProps{
//   		Input: mediaLiveInput,
//   		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type RouterInputProtocol interface {
}

// The jsii proxy struct for RouterInputProtocol
type jsiiProxy_RouterInputProtocol struct {
	_ byte // padding
}

// Create a RIST protocol configuration.
//
// Returns: RouterInputProtocol instance configured for RIST.
// Experimental.
func RouterInputProtocol_Rist(props *RistProtocolProps) RouterInputProtocol {
	_init_.Initialize()

	if err := validateRouterInputProtocol_RistParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocol",
		"rist",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an RTP protocol configuration.
//
// Returns: RouterInputProtocol instance configured for RTP.
// Experimental.
func RouterInputProtocol_Rtp(props *RtpProtocolProps) RouterInputProtocol {
	_init_.Initialize()

	if err := validateRouterInputProtocol_RtpParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocol",
		"rtp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT Caller protocol configuration.
//
// Returns: RouterInputProtocol instance configured for SRT Caller.
// Experimental.
func RouterInputProtocol_SrtCaller(props *SrtCallerProtocolProps) RouterInputProtocol {
	_init_.Initialize()

	if err := validateRouterInputProtocol_SrtCallerParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocol",
		"srtCaller",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT Listener protocol configuration.
//
// Returns: RouterInputProtocol instance configured for SRT Listener.
// Experimental.
func RouterInputProtocol_SrtListener(props *SrtListenerProtocolProps) RouterInputProtocol {
	_init_.Initialize()

	if err := validateRouterInputProtocol_SrtListenerParameters(props); err != nil {
		panic(err)
	}
	var returns RouterInputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterInputProtocol",
		"srtListener",
		[]interface{}{props},
		&returns,
	)

	return returns
}

