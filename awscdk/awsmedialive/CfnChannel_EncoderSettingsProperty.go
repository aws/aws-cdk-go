package awsmedialive


// The settings for the encoding of outputs.
//
// This entity is at the top level in the channel.
//
// Example:
//
//
type CfnChannel_EncoderSettingsProperty struct {
	// The encoding information for output audio.
	AudioDescriptions interface{} `field:"optional" json:"audioDescriptions" yaml:"audioDescriptions"`
	// The settings for ad avail blanking.
	AvailBlanking interface{} `field:"optional" json:"availBlanking" yaml:"availBlanking"`
	// The configuration settings for the ad avail handling.
	AvailConfiguration interface{} `field:"optional" json:"availConfiguration" yaml:"availConfiguration"`
	// The settings for the blackout slate.
	BlackoutSlate interface{} `field:"optional" json:"blackoutSlate" yaml:"blackoutSlate"`
	// The encoding information for output captions.
	CaptionDescriptions interface{} `field:"optional" json:"captionDescriptions" yaml:"captionDescriptions"`
	// Settings to enable specific features.
	FeatureActivations interface{} `field:"optional" json:"featureActivations" yaml:"featureActivations"`
	// The configuration settings that apply to the entire channel.
	GlobalConfiguration interface{} `field:"optional" json:"globalConfiguration" yaml:"globalConfiguration"`
	// Settings to enable and configure the motion graphics overlay feature in the channel.
	MotionGraphicsConfiguration interface{} `field:"optional" json:"motionGraphicsConfiguration" yaml:"motionGraphicsConfiguration"`
	// The settings to configure Nielsen watermarks.
	NielsenConfiguration interface{} `field:"optional" json:"nielsenConfiguration" yaml:"nielsenConfiguration"`
	// The settings for the output groups in the channel.
	OutputGroups interface{} `field:"optional" json:"outputGroups" yaml:"outputGroups"`
	// Contains settings used to acquire and adjust timecode information from the inputs.
	TimecodeConfig interface{} `field:"optional" json:"timecodeConfig" yaml:"timecodeConfig"`
	// The encoding information for output videos.
	VideoDescriptions interface{} `field:"optional" json:"videoDescriptions" yaml:"videoDescriptions"`
}

