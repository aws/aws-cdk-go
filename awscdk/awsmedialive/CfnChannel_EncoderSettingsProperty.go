package awsmedialive


// The settings for the encoding of outputs.
//
// This entity is at the top level in the channel.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html
//
type CfnChannel_EncoderSettingsProperty struct {
	// The encoding information for output audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-audiodescriptions
	//
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// The settings for ad avail blanking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availblanking
	//
	AvailBlanking interface{} `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// The configuration settings for the ad avail handling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availconfiguration
	//
	AvailConfiguration interface{} `field:"optional" json:"availConfiguration" yaml:"availConfiguration"`
	// The settings for the blackout slate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-blackoutslate
	//
	BlackoutSlate interface{} `field:"optional" json:"blackoutSlate" yaml:"blackoutSlate"`
	// The encoding information for output captions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-captiondescriptions
	//
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// Settings to enable specific features.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-featureactivations
	//
	FeatureActivations interface{} `field:"optional" json:"featureActivations" yaml:"featureActivations"`
	// The configuration settings that apply to the entire channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-globalconfiguration
	//
	GlobalConfiguration interface{} `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-motiongraphicsconfiguration
	//
	MotionGraphicsConfiguration interface{} `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// The settings to configure Nielsen watermarks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-nielsenconfiguration
	//
	NielsenConfiguration interface{} `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// The settings for the output groups in the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-outputgroups
	//
	OutputGroups interface{} `field:"optional" json:"outputGroups" yaml:"outputGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-thumbnailconfiguration
	//
	ThumbnailConfiguration interface{} `field:"optional" json:"thumbnailConfiguration" yaml:"thumbnailConfiguration"`
	// Contains settings used to acquire and adjust timecode information from the inputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-timecodeconfig
	//
	TimecodeConfig interface{} `field:"optional" json:"timecodeConfig" yaml:"timecodeConfig"`
	// The encoding information for output videos.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-videodescriptions
	//
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}

