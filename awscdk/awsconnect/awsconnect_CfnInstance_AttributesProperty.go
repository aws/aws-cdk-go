package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributesProperty := &attributesProperty{
//   	inboundCalls: jsii.Boolean(false),
//   	outboundCalls: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	autoResolveBestVoices: jsii.Boolean(false),
//   	contactflowLogs: jsii.Boolean(false),
//   	contactLens: jsii.Boolean(false),
//   	earlyMedia: jsii.Boolean(false),
//   	useCustomTtsVoices: jsii.Boolean(false),
//   }
//
type CfnInstance_AttributesProperty struct {
	// `CfnInstance.AttributesProperty.InboundCalls`.
	InboundCalls interface{} `field:"required" json:"inboundCalls" yaml:"inboundCalls"`
	// `CfnInstance.AttributesProperty.OutboundCalls`.
	OutboundCalls interface{} `field:"required" json:"outboundCalls" yaml:"outboundCalls"`
	// `CfnInstance.AttributesProperty.AutoResolveBestVoices`.
	AutoResolveBestVoices interface{} `field:"optional" json:"autoResolveBestVoices" yaml:"autoResolveBestVoices"`
	// `CfnInstance.AttributesProperty.ContactflowLogs`.
	ContactflowLogs interface{} `field:"optional" json:"contactflowLogs" yaml:"contactflowLogs"`
	// `CfnInstance.AttributesProperty.ContactLens`.
	ContactLens interface{} `field:"optional" json:"contactLens" yaml:"contactLens"`
	// `CfnInstance.AttributesProperty.EarlyMedia`.
	EarlyMedia interface{} `field:"optional" json:"earlyMedia" yaml:"earlyMedia"`
	// `CfnInstance.AttributesProperty.UseCustomTTSVoices`.
	UseCustomTtsVoices interface{} `field:"optional" json:"useCustomTtsVoices" yaml:"useCustomTtsVoices"`
}

