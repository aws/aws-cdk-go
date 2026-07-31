package awsmediaconnectalpha


// Properties for MediaConnect Flow Router Input configuration - without a connection.
//
// Example:
//   var stack Stack
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FlowInputNoConnection"), &RouterInputProps{
//   	RouterInputName: jsii.String("flow-input-no-connection"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaConnectFlowWithoutConnection(&MediaConnectFlowConfigurationWithoutConnectionProps{
//   		AvailabilityZone: jsii.String("us-east-1a"),
//   	}),
//   })
//
// Experimental.
type MediaConnectFlowConfigurationWithoutConnectionProps struct {
	// Availability Zone the router input will be placed in.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

