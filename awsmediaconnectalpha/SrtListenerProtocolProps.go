package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for SRT Listener protocol configuration.
//
// Example:
//   var stack Stack
//   var mediaLiveInput CfnInput
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
//   		MediaLiveInputArn: mediaLiveInput.attrArn,
//   		MediaLivePipelineId: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type SrtListenerProtocolProps struct {
	// Minimum latency for SRT.
	// Experimental.
	MinimumLatency awscdk.Duration `field:"required" json:"minimumLatency" yaml:"minimumLatency"`
	// Port number for SRT listener.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Optional decryption configuration for encrypted SRT streams.
	// Default: - No decryption.
	//
	// Experimental.
	DecryptionConfiguration *RouterSrtEncryption `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
}

