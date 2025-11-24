package mixinsawsconnect


// *This is a preview release for Amazon Connect .
//
// It is subject to change.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributesProperty := &AttributesProperty{
//   	AutoResolveBestVoices: jsii.Boolean(false),
//   	ContactflowLogs: jsii.Boolean(false),
//   	ContactLens: jsii.Boolean(false),
//   	EarlyMedia: jsii.Boolean(false),
//   	EnhancedChatMonitoring: jsii.Boolean(false),
//   	EnhancedContactMonitoring: jsii.Boolean(false),
//   	HighVolumeOutBound: jsii.Boolean(false),
//   	InboundCalls: jsii.Boolean(false),
//   	MultiPartyChatConference: jsii.Boolean(false),
//   	MultiPartyConference: jsii.Boolean(false),
//   	OutboundCalls: jsii.Boolean(false),
//   	UseCustomTtsVoices: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html
//
type CfnInstancePropsMixin_AttributesProperty struct {
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
	// Boolean flag which enables ENHANCED_CHAT_MONITORING on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-enhancedchatmonitoring
	//
	EnhancedChatMonitoring interface{} `field:"optional" json:"enhancedChatMonitoring" yaml:"enhancedChatMonitoring"`
	// Boolean flag which enables ENHANCED_CONTACT_MONITORING on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-enhancedcontactmonitoring
	//
	EnhancedContactMonitoring interface{} `field:"optional" json:"enhancedContactMonitoring" yaml:"enhancedContactMonitoring"`
	// Boolean flag which enables HIGH_VOLUME_OUTBOUND on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-highvolumeoutbound
	//
	HighVolumeOutBound interface{} `field:"optional" json:"highVolumeOutBound" yaml:"highVolumeOutBound"`
	// Mandatory element which enables inbound calls on new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-inboundcalls
	//
	InboundCalls interface{} `field:"optional" json:"inboundCalls" yaml:"inboundCalls"`
	// Boolean flag which enables MULTI_PARTY_CHAT_CONFERENCE on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-multipartychatconference
	//
	MultiPartyChatConference interface{} `field:"optional" json:"multiPartyChatConference" yaml:"multiPartyChatConference"`
	// Boolean flag which enables MULTI_PARTY_CONFERENCE on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-multipartyconference
	//
	MultiPartyConference interface{} `field:"optional" json:"multiPartyConference" yaml:"multiPartyConference"`
	// Mandatory element which enables outbound calls on new instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-outboundcalls
	//
	OutboundCalls interface{} `field:"optional" json:"outboundCalls" yaml:"outboundCalls"`
	// Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-instance-attributes.html#cfn-connect-instance-attributes-usecustomttsvoices
	//
	UseCustomTtsVoices interface{} `field:"optional" json:"useCustomTtsVoices" yaml:"useCustomTtsVoices"`
}

