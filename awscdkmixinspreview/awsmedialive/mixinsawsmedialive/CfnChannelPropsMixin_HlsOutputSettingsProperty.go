package mixinsawsmedialive


// The settings for an HLS output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsOutputSettingsProperty := &HlsOutputSettingsProperty{
//   	H265PackagingType: jsii.String("h265PackagingType"),
//   	HlsSettings: &HlsSettingsProperty{
//   		AudioOnlyHlsSettings: &AudioOnlyHlsSettingsProperty{
//   			AudioGroupId: jsii.String("audioGroupId"),
//   			AudioOnlyImage: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			AudioTrackType: jsii.String("audioTrackType"),
//   			SegmentType: jsii.String("segmentType"),
//   		},
//   		Fmp4HlsSettings: &Fmp4HlsSettingsProperty{
//   			AudioRenditionSets: jsii.String("audioRenditionSets"),
//   			NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		},
//   		FrameCaptureHlsSettings: &FrameCaptureHlsSettingsProperty{
//   		},
//   		StandardHlsSettings: &StandardHlsSettingsProperty{
//   			AudioRenditionSets: jsii.String("audioRenditionSets"),
//   			M3U8Settings: &M3u8SettingsProperty{
//   				AudioFramesPerPes: jsii.Number(123),
//   				AudioPids: jsii.String("audioPids"),
//   				EcmPid: jsii.String("ecmPid"),
//   				KlvBehavior: jsii.String("klvBehavior"),
//   				KlvDataPids: jsii.String("klvDataPids"),
//   				NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				PatInterval: jsii.Number(123),
//   				PcrControl: jsii.String("pcrControl"),
//   				PcrPeriod: jsii.Number(123),
//   				PcrPid: jsii.String("pcrPid"),
//   				PmtInterval: jsii.Number(123),
//   				PmtPid: jsii.String("pmtPid"),
//   				ProgramNum: jsii.Number(123),
//   				Scte35Behavior: jsii.String("scte35Behavior"),
//   				Scte35Pid: jsii.String("scte35Pid"),
//   				TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				TimedMetadataPid: jsii.String("timedMetadataPid"),
//   				TransportStreamId: jsii.Number(123),
//   				VideoPid: jsii.String("videoPid"),
//   			},
//   		},
//   	},
//   	NameModifier: jsii.String("nameModifier"),
//   	SegmentModifier: jsii.String("segmentModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsoutputsettings.html
//
type CfnChannelPropsMixin_HlsOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsoutputsettings.html#cfn-medialive-channel-hlsoutputsettings-h265packagingtype
	//
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// The settings regarding the underlying stream.
	//
	// These settings are different for audio-only outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsoutputsettings.html#cfn-medialive-channel-hlsoutputsettings-hlssettings
	//
	HlsSettings interface{} `field:"optional" json:"hlsSettings" yaml:"hlsSettings"`
	// A string that is concatenated to the end of the destination file name.
	//
	// Accepts \"Format Identifiers\":#formatIdentifierParameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsoutputsettings.html#cfn-medialive-channel-hlsoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// A string that is concatenated to the end of segment file names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsoutputsettings.html#cfn-medialive-channel-hlsoutputsettings-segmentmodifier
	//
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

