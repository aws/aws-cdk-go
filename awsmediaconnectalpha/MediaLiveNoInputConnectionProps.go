package awsmediaconnectalpha


// Properties for MediaLive Router Output configuration without specific input connection.
//
// Example:
//   var stack Stack
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("MediaLiveOutputNoConnection"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("medialive-output-no-connection"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(15)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_GLOBAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInputWithoutConnection(&MediaLiveNoInputConnectionProps{
//   		AvailabilityZone: jsii.String("us-east-1a"),
//   	}),
//   })
//
// Experimental.
type MediaLiveNoInputConnectionProps struct {
	// Availability zone for the router output when not connecting to a specific input.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	DestinationTransitEncryption *TransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
}

