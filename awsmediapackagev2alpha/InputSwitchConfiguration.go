package awsmediapackagev2alpha


// Input Switch Configuration.
//
// Example:
//   var stack Stack
//   var group ChannelGroup
//
//
//   hlsChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("HlsChannel"), &ChannelProps{
//   	ChannelGroup: group,
//   	Input: awsmediapackagev2alpha.InputConfiguration_Hls(),
//   })
//
//   cmafChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("CmafChannel"), &ChannelProps{
//   	ChannelGroup: group,
//   	Input: awsmediapackagev2alpha.InputConfiguration_Cmaf(&CmafInputProps{
//   		InputSwitchConfiguration: &InputSwitchConfiguration{
//   			MqcsInputSwitching: jsii.Boolean(true),
//   		},
//   		OutputHeaders: []HeadersCMSD{
//   			awsmediapackagev2alpha.HeadersCMSD_MQCS(),
//   		},
//   	}),
//   })
//
//   simpleCmafChannel := awsmediapackagev2alpha.NewChannel(stack, jsii.String("SimpleCmafChannel"), &ChannelProps{
//   	ChannelGroup: group,
//   	Input: awsmediapackagev2alpha.InputConfiguration_*Cmaf(&CmafInputProps{
//   		OutputHeaders: []HeadersCMSD{
//   			awsmediapackagev2alpha.HeadersCMSD_MQCS(),
//   		},
//   	}),
//   })
//
// Experimental.
type InputSwitchConfiguration struct {
	// When true, AWS Elemental MediaPackage performs input switching based on the MQCS.
	//
	// This setting is valid only when InputType is CMAF.
	// Default: false.
	//
	// Experimental.
	MqcsInputSwitching *bool `field:"optional" json:"mqcsInputSwitching" yaml:"mqcsInputSwitching"`
	// For CMAF inputs, indicates which input MediaPackage should prefer when both inputs have equal MQCS scores.
	//
	// Select 1 to prefer the first ingest endpoint, or 2 to prefer the second ingest endpoint.
	// If you don't specify a preferred input, MediaPackage uses its default switching behavior when MQCS scores are equal.
	// Default: - MediaPackage uses its default switching behavior.
	//
	// Experimental.
	PreferredInput IngestEndpoint `field:"optional" json:"preferredInput" yaml:"preferredInput"`
}

