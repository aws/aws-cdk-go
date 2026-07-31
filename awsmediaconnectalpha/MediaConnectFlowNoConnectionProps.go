package awsmediaconnectalpha


// Properties for MediaConnect Flow Router Output configuration without specific flow connection.
//
// Example:
//   var stack Stack
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("FlowOutputNoConnection"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("flow-output-no-connection"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_100(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlowWithoutConnection(&MediaConnectFlowNoConnectionProps{
//   		AvailabilityZone: jsii.String("us-east-1a"),
//   	}),
//   })
//
// Experimental.
type MediaConnectFlowNoConnectionProps struct {
	// Availability zone for the router output when not connecting to a specific flow.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	DestinationTransitEncryption *TransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
}

