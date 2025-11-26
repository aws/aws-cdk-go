package previewawsmedialivemixins


// The configuration of a UDP output group.
//
// The parent of this entity is OutputGroupSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   udpGroupSettingsProperty := &UdpGroupSettingsProperty{
//   	InputLossAction: jsii.String("inputLossAction"),
//   	TimedMetadataId3Frame: jsii.String("timedMetadataId3Frame"),
//   	TimedMetadataId3Period: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html
//
type CfnChannelPropsMixin_UdpGroupSettingsProperty struct {
	// Specifies the behavior of the last resort when the input video is lost, and no more backup inputs are available.
	//
	// When dropTs is selected, the entire transport stream stops emitting. When dropProgram is selected, the program can be dropped from the transport stream (and replaced with null packets to meet the TS bitrate requirement). Or when emitProgram is selected, the transport stream continues to be produced normally with repeat frames, black frames, or slate frames substituted for the absent input video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-inputlossaction
	//
	InputLossAction *string `field:"optional" json:"inputLossAction" yaml:"inputLossAction"`
	// Indicates the ID3 frame that has the timecode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-timedmetadataid3frame
	//
	TimedMetadataId3Frame *string `field:"optional" json:"timedMetadataId3Frame" yaml:"timedMetadataId3Frame"`
	// The timed metadata interval in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-udpgroupsettings.html#cfn-medialive-channel-udpgroupsettings-timedmetadataid3period
	//
	TimedMetadataId3Period *float64 `field:"optional" json:"timedMetadataId3Period" yaml:"timedMetadataId3Period"`
}

