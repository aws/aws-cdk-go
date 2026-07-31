package awsmediaconnectalpha


// Properties for MediaLive Router Output configuration with specific input connection.
//
// Example:
//   var stack Stack
//   var mediaLiveInput CfnInput
//
//
//   output := awsmediaconnectalpha.NewRouterOutput(stack, jsii.String("MediaLiveOutput"), &RouterOutputProps{
//   	RouterOutputName: jsii.String("medialive-output"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(15)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_GLOBAL(),
//   	Tier: awsmediaconnectalpha.RouterOutputTier_OUTPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterOutputConfiguration_MediaLiveInput(&MediaLiveInputConnectionProps{
//   		MediaLiveInputArn: mediaLiveInput.attrArn,
//   		MediaLivePipelineId: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   	}),
//   })
//
// Experimental.
type MediaLiveInputConnectionProps struct {
	// ARN of the MediaLive input to send output to.
	//
	// Note: This will change to accept an IInputRef (typed MediaLive Input reference)
	// when the.
	// Experimental.
	MediaLiveInputArn *string `field:"required" json:"mediaLiveInputArn" yaml:"mediaLiveInputArn"`
	// Pipeline ID for MediaLive input.
	// Experimental.
	MediaLivePipelineId MediaLivePipeline `field:"required" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	DestinationTransitEncryption *TransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
}

