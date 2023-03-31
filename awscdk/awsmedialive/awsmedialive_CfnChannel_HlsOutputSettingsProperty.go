package awsmedialive


// The settings for an HLS output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsOutputSettingsProperty := &hlsOutputSettingsProperty{
//   	h265PackagingType: jsii.String("h265PackagingType"),
//   	hlsSettings: &hlsSettingsProperty{
//   		audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   			audioGroupId: jsii.String("audioGroupId"),
//   			audioOnlyImage: &inputLocationProperty{
//   				passwordParam: jsii.String("passwordParam"),
//   				uri: jsii.String("uri"),
//   				username: jsii.String("username"),
//   			},
//   			audioTrackType: jsii.String("audioTrackType"),
//   			segmentType: jsii.String("segmentType"),
//   		},
//   		fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   			audioRenditionSets: jsii.String("audioRenditionSets"),
//   			nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		},
//   		frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   		},
//   		standardHlsSettings: &standardHlsSettingsProperty{
//   			audioRenditionSets: jsii.String("audioRenditionSets"),
//   			m3U8Settings: &m3u8SettingsProperty{
//   				audioFramesPerPes: jsii.Number(123),
//   				audioPids: jsii.String("audioPids"),
//   				ecmPid: jsii.String("ecmPid"),
//   				nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				patInterval: jsii.Number(123),
//   				pcrControl: jsii.String("pcrControl"),
//   				pcrPeriod: jsii.Number(123),
//   				pcrPid: jsii.String("pcrPid"),
//   				pmtInterval: jsii.Number(123),
//   				pmtPid: jsii.String("pmtPid"),
//   				programNum: jsii.Number(123),
//   				scte35Behavior: jsii.String("scte35Behavior"),
//   				scte35Pid: jsii.String("scte35Pid"),
//   				timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				timedMetadataPid: jsii.String("timedMetadataPid"),
//   				transportStreamId: jsii.Number(123),
//   				videoPid: jsii.String("videoPid"),
//   			},
//   		},
//   	},
//   	nameModifier: jsii.String("nameModifier"),
//   	segmentModifier: jsii.String("segmentModifier"),
//   }
//
type CfnChannel_HlsOutputSettingsProperty struct {
	// Only applicable when this output is referencing an H.265 video description. Specifies whether MP4 segments should be packaged as HEV1 or HVC1.
	H265PackagingType *string `field:"optional" json:"h265PackagingType" yaml:"h265PackagingType"`
	// The settings regarding the underlying stream.
	//
	// These settings are different for audio-only outputs.
	HlsSettings interface{} `field:"optional" json:"hlsSettings" yaml:"hlsSettings"`
	// A string that is concatenated to the end of the destination file name.
	//
	// Accepts \"Format Identifiers\":#formatIdentifierParameters.
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
	// A string that is concatenated to the end of segment file names.
	SegmentModifier *string `field:"optional" json:"segmentModifier" yaml:"segmentModifier"`
}

