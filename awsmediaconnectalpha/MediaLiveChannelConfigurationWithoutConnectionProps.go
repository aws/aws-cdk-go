package awsmediaconnectalpha


// Properties for MediaLive Channel Router Input configuration without a specific channel connection.
//
// Use this when you want to set up the router input before the target MediaLive channel exists.
//
// Example:
//   var stack Stack
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInputNoConnection"), &RouterInputProps{
//   	RouterInputName: jsii.String("channel-input-no-connection"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannelWithoutConnection(&MediaLiveChannelConfigurationWithoutConnectionProps{
//   		AvailabilityZone: jsii.String("us-east-1a"),
//   	}),
//   })
//
// Experimental.
type MediaLiveChannelConfigurationWithoutConnectionProps struct {
	// Availability Zone the router input will be placed in.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

