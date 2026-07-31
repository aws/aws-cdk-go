package awsmediaconnectalpha


// MediaLive pipeline options.
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
type MediaLivePipeline string

const (
	// Pipeline 0.
	// Experimental.
	MediaLivePipeline_PIPELINE_0 MediaLivePipeline = "PIPELINE_0"
	// Pipeline 1.
	// Experimental.
	MediaLivePipeline_PIPELINE_1 MediaLivePipeline = "PIPELINE_1"
)

