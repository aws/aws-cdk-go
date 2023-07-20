package awsmedialive


// Settings for the fMP4 containers.
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fmp4HlsSettingsProperty := &Fmp4HlsSettingsProperty{
//   	AudioRenditionSets: jsii.String("audioRenditionSets"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fmp4hlssettings.html
//
type CfnChannel_Fmp4HlsSettingsProperty struct {
	// List all the audio groups that are used with the video output stream.
	//
	// Input all the audio GROUP-IDs that are associated to the video, separate by ','.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fmp4hlssettings.html#cfn-medialive-channel-fmp4hlssettings-audiorenditionsets
	//
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fmp4hlssettings.html#cfn-medialive-channel-fmp4hlssettings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// When set to passthrough, timed metadata is passed through from input to output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-fmp4hlssettings.html#cfn-medialive-channel-fmp4hlssettings-timedmetadatabehavior
	//
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
}

