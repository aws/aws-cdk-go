package awsmedialive


// The configuration of an HLS output that is a standard output (not an audio-only output).
//
// The parent of this entity is HlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardHlsSettingsProperty := &StandardHlsSettingsProperty{
//   	AudioRenditionSets: jsii.String("audioRenditionSets"),
//   	M3U8Settings: &M3u8SettingsProperty{
//   		AudioFramesPerPes: jsii.Number(123),
//   		AudioPids: jsii.String("audioPids"),
//   		EcmPid: jsii.String("ecmPid"),
//   		NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		PatInterval: jsii.Number(123),
//   		PcrControl: jsii.String("pcrControl"),
//   		PcrPeriod: jsii.Number(123),
//   		PcrPid: jsii.String("pcrPid"),
//   		PmtInterval: jsii.Number(123),
//   		PmtPid: jsii.String("pmtPid"),
//   		ProgramNum: jsii.Number(123),
//   		Scte35Behavior: jsii.String("scte35Behavior"),
//   		Scte35Pid: jsii.String("scte35Pid"),
//   		TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		TimedMetadataPid: jsii.String("timedMetadataPid"),
//   		TransportStreamId: jsii.Number(123),
//   		VideoPid: jsii.String("videoPid"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-standardhlssettings.html
//
type CfnChannel_StandardHlsSettingsProperty struct {
	// Lists all the audio groups that are used with the video output stream.
	//
	// This inputs all the audio GROUP-IDs that are associated with the video, separated by a comma (,).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-standardhlssettings.html#cfn-medialive-channel-standardhlssettings-audiorenditionsets
	//
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Settings for the M3U8 container.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-standardhlssettings.html#cfn-medialive-channel-standardhlssettings-m3u8settings
	//
	M3U8Settings interface{} `field:"optional" json:"m3U8Settings" yaml:"m3U8Settings"`
}

