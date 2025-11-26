package previewawsmedialivemixins


// The settings for one output group.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html
//
type CfnChannelPropsMixin_OutputGroupProperty struct {
	// A custom output group name that you can optionally define.
	//
	// Only letters, numbers, and the underscore character are allowed. The maximum length is 32 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings associated with the output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputgroupsettings
	//
	OutputGroupSettings interface{} `field:"optional" json:"outputGroupSettings" yaml:"outputGroupSettings"`
	// The settings for the outputs in the output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputgroup.html#cfn-medialive-channel-outputgroup-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

