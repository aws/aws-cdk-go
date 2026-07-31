package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for SRT Listener protocol configuration for outputs.
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
type SrtListenerOutputProtocolProps struct {
	// Minimum latency for SRT.
	// Experimental.
	MinimumLatency awscdk.Duration `field:"required" json:"minimumLatency" yaml:"minimumLatency"`
	// Port number for SRT listener.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Optional encryption configuration for SRT streams.
	// Default: - No encryption.
	//
	// Experimental.
	EncryptionConfiguration *RouterSrtEncryption `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
}

