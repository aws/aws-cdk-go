package previewawsmedialivemixins


// Properties for CfnChannelPropsMixin.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html
//
type CfnChannelMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-anywheresettings
	//
	AnywhereSettings interface{} `field:"optional" json:"anywhereSettings" yaml:"anywhereSettings"`
	// Specification of CDI inputs for this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-cdiinputspecification
	//
	CdiInputSpecification interface{} `field:"optional" json:"cdiInputSpecification" yaml:"cdiInputSpecification"`
	// The class for this channel.
	//
	// For a channel with two pipelines, the class is STANDARD. For a channel with one pipeline, the class is SINGLE_PIPELINE.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-channelclass
	//
	ChannelClass *string `field:"optional" json:"channelClass" yaml:"channelClass"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-channelengineversion
	//
	ChannelEngineVersion interface{} `field:"optional" json:"channelEngineVersion" yaml:"channelEngineVersion"`
	// The settings that identify the destination for the outputs in this MediaLive output package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-destinations
	//
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-dryrun
	//
	DryRun interface{} `field:"optional" json:"dryRun" yaml:"dryRun"`
	// The encoding configuration for the output content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-encodersettings
	//
	EncoderSettings interface{} `field:"optional" json:"encoderSettings" yaml:"encoderSettings"`
	// The list of input attachments for the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputattachments
	//
	InputAttachments interface{} `field:"optional" json:"inputAttachments" yaml:"inputAttachments"`
	// The input specification for this channel.
	//
	// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-inputspecification
	//
	InputSpecification interface{} `field:"optional" json:"inputSpecification" yaml:"inputSpecification"`
	// The verbosity for logging activity for this channel.
	//
	// Charges for logging (which are generated through Amazon CloudWatch Logging) are higher for higher verbosities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-loglevel
	//
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Maintenance settings for this channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-maintenance
	//
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Name of channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when running this channel.
	//
	// The role is identified by its ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A collection of tags for this channel.
	//
	// Each tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html#cfn-medialive-channel-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

