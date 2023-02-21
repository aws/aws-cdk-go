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
//   hlsSettingsProperty := &hlsSettingsProperty{
//   	audioOnlyHlsSettings: &audioOnlyHlsSettingsProperty{
//   		audioGroupId: jsii.String("audioGroupId"),
//   		audioOnlyImage: &inputLocationProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			uri: jsii.String("uri"),
//   			username: jsii.String("username"),
//   		},
//   		audioTrackType: jsii.String("audioTrackType"),
//   		segmentType: jsii.String("segmentType"),
//   	},
//   	fmp4HlsSettings: &fmp4HlsSettingsProperty{
//   		audioRenditionSets: jsii.String("audioRenditionSets"),
//   		nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	},
//   	frameCaptureHlsSettings: &frameCaptureHlsSettingsProperty{
//   	},
//   	standardHlsSettings: &standardHlsSettingsProperty{
//   		audioRenditionSets: jsii.String("audioRenditionSets"),
//   		m3U8Settings: &m3u8SettingsProperty{
//   			audioFramesPerPes: jsii.Number(123),
//   			audioPids: jsii.String("audioPids"),
//   			ecmPid: jsii.String("ecmPid"),
//   			nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			patInterval: jsii.Number(123),
//   			pcrControl: jsii.String("pcrControl"),
//   			pcrPeriod: jsii.Number(123),
//   			pcrPid: jsii.String("pcrPid"),
//   			pmtInterval: jsii.Number(123),
//   			pmtPid: jsii.String("pmtPid"),
//   			programNum: jsii.Number(123),
//   			scte35Behavior: jsii.String("scte35Behavior"),
//   			scte35Pid: jsii.String("scte35Pid"),
//   			timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			timedMetadataPid: jsii.String("timedMetadataPid"),
//   			transportStreamId: jsii.Number(123),
//   			videoPid: jsii.String("videoPid"),
//   		},
//   	},
//   }
//
type CfnChannel_HlsSettingsProperty struct {
	// The settings for an audio-only output.
	AudioOnlyHlsSettings interface{} `field:"optional" json:"audioOnlyHlsSettings" yaml:"audioOnlyHlsSettings"`
	// The settings for an fMP4 container.
	Fmp4HlsSettings interface{} `field:"optional" json:"fmp4HlsSettings" yaml:"fmp4HlsSettings"`
	// Settings for a frame capture output in an HLS output group.
	FrameCaptureHlsSettings interface{} `field:"optional" json:"frameCaptureHlsSettings" yaml:"frameCaptureHlsSettings"`
	// The settings for a standard output (an output that is not audio-only).
	StandardHlsSettings interface{} `field:"optional" json:"standardHlsSettings" yaml:"standardHlsSettings"`
}

