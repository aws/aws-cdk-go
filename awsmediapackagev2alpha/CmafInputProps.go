package awsmediapackagev2alpha


// Properties for CMAF input configuration.
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
type CmafInputProps struct {
	// The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.
	// Default: No customized input switch configuration added.
	//
	// Experimental.
	InputSwitchConfiguration *InputSwitchConfiguration `field:"optional" json:"inputSwitchConfiguration" yaml:"inputSwitchConfiguration"`
	// The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.
	// Default: none.
	//
	// Experimental.
	OutputHeaders *[]HeadersCMSD `field:"optional" json:"outputHeaders" yaml:"outputHeaders"`
}

