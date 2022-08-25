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
//   standardHlsSettingsProperty := &standardHlsSettingsProperty{
//   	audioRenditionSets: jsii.String("audioRenditionSets"),
//   	m3U8Settings: &m3u8SettingsProperty{
//   		audioFramesPerPes: jsii.Number(123),
//   		audioPids: jsii.String("audioPids"),
//   		ecmPid: jsii.String("ecmPid"),
//   		nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		patInterval: jsii.Number(123),
//   		pcrControl: jsii.String("pcrControl"),
//   		pcrPeriod: jsii.Number(123),
//   		pcrPid: jsii.String("pcrPid"),
//   		pmtInterval: jsii.Number(123),
//   		pmtPid: jsii.String("pmtPid"),
//   		programNum: jsii.Number(123),
//   		scte35Behavior: jsii.String("scte35Behavior"),
//   		scte35Pid: jsii.String("scte35Pid"),
//   		timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		timedMetadataPid: jsii.String("timedMetadataPid"),
//   		transportStreamId: jsii.Number(123),
//   		videoPid: jsii.String("videoPid"),
//   	},
//   }
//
type CfnChannel_StandardHlsSettingsProperty struct {
	// Lists all the audio groups that are used with the video output stream.
	//
	// This inputs all the audio GROUP-IDs that are associated with the video, separated by a comma (,).
	AudioRenditionSets *string `field:"optional" json:"audioRenditionSets" yaml:"audioRenditionSets"`
	// Settings for the M3U8 container.
	M3U8Settings interface{} `field:"optional" json:"m3U8Settings" yaml:"m3U8Settings"`
}

