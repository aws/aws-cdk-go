package awsconnect


// *This is a preview release for Amazon Connect .
//
// It is subject to change.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributesProperty := &AttributesProperty{
//   	InboundCalls: jsii.Boolean(false),
//   	OutboundCalls: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AutoResolveBestVoices: jsii.Boolean(false),
//   	ContactflowLogs: jsii.Boolean(false),
//   	ContactLens: jsii.Boolean(false),
//   	EarlyMedia: jsii.Boolean(false),
//   	UseCustomTtsVoices: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html
//
type CfnInstance_AttributesProperty struct {
	// Mandatory element which enables inbound calls on new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-inboundcalls
	//
	InboundCalls interface{} `field:"required" json:"inboundCalls" yaml:"inboundCalls"`
	// Mandatory element which enables outbound calls on new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-outboundcalls
	//
	OutboundCalls interface{} `field:"required" json:"outboundCalls" yaml:"outboundCalls"`
	// Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-autoresolvebestvoices
	//
	AutoResolveBestVoices interface{} `field:"optional" json:"autoResolveBestVoices" yaml:"autoResolveBestVoices"`
	// Boolean flag which enables CONTACTFLOW_LOGS on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-contactflowlogs
	//
	ContactflowLogs interface{} `field:"optional" json:"contactflowLogs" yaml:"contactflowLogs"`
	// Boolean flag which enables CONTACT_LENS on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-contactlens
	//
	ContactLens interface{} `field:"optional" json:"contactLens" yaml:"contactLens"`
	// Boolean flag which enables EARLY_MEDIA on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-earlymedia
	//
	EarlyMedia interface{} `field:"optional" json:"earlyMedia" yaml:"earlyMedia"`
	// Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-usecustomttsvoices
	//
	UseCustomTtsVoices interface{} `field:"optional" json:"useCustomTtsVoices" yaml:"useCustomTtsVoices"`
}

