package awsmediaconnectalpha


// Properties for MediaConnect Flow Router Input configuration.
//
// Example:
//   var stack Stack
//   var flow Flow
//   var flowOutput FlowOutput
//
//
//   input := awsmediaconnectalpha.NewRouterInput(stack, jsii.String("FlowInput"), &RouterInputProps{
//   	RouterInputName: jsii.String("flow-input"),
//   	MaximumBitrate: awscdk.Bitrate_Mbps(jsii.Number(20)),
//   	RoutingScope: awsmediaconnectalpha.RoutingScope_REGIONAL(),
//   	Tier: awsmediaconnectalpha.RouterInputTier_INPUT_50(),
//   	Configuration: awsmediaconnectalpha.RouterInputConfiguration_MediaConnectFlow(&MediaConnectFlowConfigurationProps{
//   		Flow: flow,
//   		FlowOutput: flowOutput,
//   	}),
//   })
//
// Experimental.
type MediaConnectFlowConfigurationProps struct {
	// The MediaConnect flow to use as input.
	// Experimental.
	Flow IFlow `field:"required" json:"flow" yaml:"flow"`
	// The flow output that feeds this router input.
	// Experimental.
	FlowOutput IFlowOutput `field:"required" json:"flowOutput" yaml:"flowOutput"`
	// Optional transit encryption configuration.
	// Default: - Automatic encryption will be used.
	//
	// Experimental.
	SourceTransitDecryption *TransitEncryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

