package awsmedialive


// The settings for an HLS output.
//
// The parent of this entity is HlsOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsSettingsProperty := &HlsSettingsProperty{
//   	AudioOnlyHlsSettings: &AudioOnlyHlsSettingsProperty{
//   		AudioGroupId: jsii.String("audioGroupId"),
//   		AudioOnlyImage: &InputLocationProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			Uri: jsii.String("uri"),
//   			Username: jsii.String("username"),
//   		},
//   		AudioTrackType: jsii.String("audioTrackType"),
//   		SegmentType: jsii.String("segmentType"),
//   	},
//   	Fmp4HlsSettings: &Fmp4HlsSettingsProperty{
//   		AudioRenditionSets: jsii.String("audioRenditionSets"),
//   		NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	},
//   	FrameCaptureHlsSettings: &FrameCaptureHlsSettingsProperty{
//   	},
//   	StandardHlsSettings: &StandardHlsSettingsProperty{
//   		AudioRenditionSets: jsii.String("audioRenditionSets"),
//   		M3U8Settings: &M3u8SettingsProperty{
//   			AudioFramesPerPes: jsii.Number(123),
//   			AudioPids: jsii.String("audioPids"),
//   			EcmPid: jsii.String("ecmPid"),
//   			KlvBehavior: jsii.String("klvBehavior"),
//   			KlvDataPids: jsii.String("klvDataPids"),
//   			NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			PatInterval: jsii.Number(123),
//   			PcrControl: jsii.String("pcrControl"),
//   			PcrPeriod: jsii.Number(123),
//   			PcrPid: jsii.String("pcrPid"),
//   			PmtInterval: jsii.Number(123),
//   			PmtPid: jsii.String("pmtPid"),
//   			ProgramNum: jsii.Number(123),
//   			Scte35Behavior: jsii.String("scte35Behavior"),
//   			Scte35Pid: jsii.String("scte35Pid"),
//   			TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			TimedMetadataPid: jsii.String("timedMetadataPid"),
//   			TransportStreamId: jsii.Number(123),
//   			VideoPid: jsii.String("videoPid"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html
//
type CfnChannel_HlsSettingsProperty struct {
	// The settings for an audio-only output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-audioonlyhlssettings
	//
	AudioOnlyHlsSettings interface{} `field:"optional" json:"audioOnlyHlsSettings" yaml:"audioOnlyHlsSettings"`
	// The settings for an fMP4 container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-fmp4hlssettings
	//
	Fmp4HlsSettings interface{} `field:"optional" json:"fmp4HlsSettings" yaml:"fmp4HlsSettings"`
	// Settings for a frame capture output in an HLS output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-framecapturehlssettings
	//
	FrameCaptureHlsSettings interface{} `field:"optional" json:"frameCaptureHlsSettings" yaml:"frameCaptureHlsSettings"`
	// The settings for a standard output (an output that is not audio-only).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-standardhlssettings
	//
	StandardHlsSettings interface{} `field:"optional" json:"standardHlsSettings" yaml:"standardHlsSettings"`
}

