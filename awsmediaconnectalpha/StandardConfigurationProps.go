package awsmediaconnectalpha


// Properties for standard Router Input configuration.
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
type StandardConfigurationProps struct {
	// Network interface for the Router Input.
	// Experimental.
	NetworkInterface IRouterNetworkInterface `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// Protocol configuration for the input.
	// Experimental.
	Protocol RouterInputProtocol `field:"required" json:"protocol" yaml:"protocol"`
}

