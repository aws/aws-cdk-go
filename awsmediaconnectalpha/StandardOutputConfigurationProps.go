package awsmediaconnectalpha


// Properties for standard Router Output configuration.
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
type StandardOutputConfigurationProps struct {
	// Network interface for the Router Output.
	// Experimental.
	NetworkInterface IRouterNetworkInterface `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// Protocol configuration for the output.
	// Experimental.
	Protocol RouterOutputProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// The availability zone where the router output is located.
	// Default: - assigned by the MediaConnect service.
	//
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
}

