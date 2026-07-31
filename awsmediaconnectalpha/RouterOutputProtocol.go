package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Router Output protocol configurations.
//
// Supported protocols: RTP, RIST, SRT (Listener and Caller).
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
type RouterOutputProtocol interface {
}

// The jsii proxy struct for RouterOutputProtocol
type jsiiProxy_RouterOutputProtocol struct {
	_ byte // padding
}

// Create a RIST protocol configuration.
//
// Returns: RouterOutputProtocol instance configured for RIST.
//
// Example:
//   awsmediaconnectalpha.RouterOutputProtocol_Rist(&RistOutputProtocolProps{
//   	DestinationAddress: jsii.String("10.0.0.1"),
//   	Port: jsii.Number(5000),
//   })
//
// Experimental.
func RouterOutputProtocol_Rist(props *RistOutputProtocolProps) RouterOutputProtocol {
	_init_.Initialize()

	if err := validateRouterOutputProtocol_RistParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocol",
		"rist",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an RTP protocol configuration.
//
// Returns: RouterOutputProtocol instance configured for RTP.
//
// Example:
//   awsmediaconnectalpha.RouterOutputProtocol_Rtp(&RtpOutputProtocolProps{
//   	DestinationAddress: jsii.String("10.0.0.1"),
//   	Port: jsii.Number(5000),
//   })
//
// Experimental.
func RouterOutputProtocol_Rtp(props *RtpOutputProtocolProps) RouterOutputProtocol {
	_init_.Initialize()

	if err := validateRouterOutputProtocol_RtpParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocol",
		"rtp",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT Caller protocol configuration.
//
// Returns: RouterOutputProtocol instance configured for SRT Caller.
//
// Example:
//   awsmediaconnectalpha.RouterOutputProtocol_SrtCaller(&SrtCallerOutputProtocolProps{
//   	DestinationAddress: jsii.String("10.0.0.1"),
//   	DestinationPort: jsii.Number(5000),
//   	MinimumLatency: awscdk.Duration_Millis(jsii.Number(125)),
//   })
//
// Experimental.
func RouterOutputProtocol_SrtCaller(props *SrtCallerOutputProtocolProps) RouterOutputProtocol {
	_init_.Initialize()

	if err := validateRouterOutputProtocol_SrtCallerParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocol",
		"srtCaller",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create an SRT Listener protocol configuration.
//
// Returns: RouterOutputProtocol instance configured for SRT Listener.
//
// Example:
//   awsmediaconnectalpha.RouterOutputProtocol_SrtListener(&SrtListenerOutputProtocolProps{
//   	Port: jsii.Number(5000),
//   	MinimumLatency: awscdk.Duration_Millis(jsii.Number(125)),
//   })
//
// Experimental.
func RouterOutputProtocol_SrtListener(props *SrtListenerOutputProtocolProps) RouterOutputProtocol {
	_init_.Initialize()

	if err := validateRouterOutputProtocol_SrtListenerParameters(props); err != nil {
		panic(err)
	}
	var returns RouterOutputProtocol

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.RouterOutputProtocol",
		"srtListener",
		[]interface{}{props},
		&returns,
	)

	return returns
}

