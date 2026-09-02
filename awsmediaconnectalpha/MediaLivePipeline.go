package awsmediaconnectalpha


// MediaLive pipeline options.
//
// Example:
//   var stack Stack
//   var mediaLiveChannel IChannel
//   var transitSecret Secret
//   // must hold the same value as the channel's MediaConnectRouterSettings.shared() secret
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("ChannelInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("channel-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaLiveChannel(&MediaLiveChannelConfigurationProps{
//   		Channel: mediaLiveChannel,
//   		OutputName: jsii.String("router-ts"),
//   		Pipeline: awsmediaconnectalpha.MediaLivePipeline_PIPELINE_0,
//   		SourceTransitDecryption: &TransitEncryption{
//   			Secret: transitSecret,
//   		},
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

