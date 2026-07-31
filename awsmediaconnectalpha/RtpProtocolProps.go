package awsmediaconnectalpha


// Properties for RTP protocol configuration.
//
// Example:
//   var stack Stack
//   var networkInterface RouterNetworkInterface
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("RtpInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("rtp-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(10)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	// tier defaults to RouterInputTier.INPUT_20 (lowest cost)
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_Standard(&StandardConfigurationProps{
//   		NetworkInterface: networkInterface,
//   		Protocol: awsmediaconnectalpha.RouterInputProtocol_Rtp(&RtpProtocolProps{
//   			Port: jsii.Number(5000),
//   		}),
//   	}),
//   })
//
// Experimental.
type RtpProtocolProps struct {
	// Port number for RTP traffic.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Forward Error Correction setting.
	// Default: - undefined; when omitted, MediaConnect applies ForwardErrorCorrection.DISABLED at deploy time
	//
	// Experimental.
	ForwardErrorCorrection ForwardErrorCorrection `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
}

