package awsmediaconnectalpha


// Properties for MediaConnect Flow Router Output configuration with specific flow connection.
//
// Example:
//   var stack Stack
//   var flow Flow
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("FlowOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("flow-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_100(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaConnectFlow(&MediaConnectFlowConnectionProps{
//   		Flow: flow,
//   	}),
//   })
//
// Experimental.
type MediaConnectFlowConnectionProps struct {
	// MediaConnect Flow to send output to.
	// Experimental.
	Flow IFlow `field:"required" json:"flow" yaml:"flow"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	DestinationTransitEncryption *TransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
}

