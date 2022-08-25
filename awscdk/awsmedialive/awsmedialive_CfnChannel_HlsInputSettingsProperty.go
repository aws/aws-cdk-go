package awsmedialive


// Information about how to connect to the upstream system.
//
// The parent of this entity is NetworkInputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsInputSettingsProperty := &hlsInputSettingsProperty{
//   	bandwidth: jsii.Number(123),
//   	bufferSegments: jsii.Number(123),
//   	retries: jsii.Number(123),
//   	retryInterval: jsii.Number(123),
//   	scte35Source: jsii.String("scte35Source"),
//   }
//
type CfnChannel_HlsInputSettingsProperty struct {
	// When specified, the HLS stream with the m3u8 bandwidth that most closely matches this value is chosen.
	//
	// Otherwise, the highest bandwidth stream in the m3u8 is chosen. The bitrate is specified in bits per second, as in an HLS manifest.
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// When specified, reading of the HLS input begins this many buffer segments from the end (most recently written segment).
	//
	// When not specified, the HLS input begins with the first segment specified in the m3u8.
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// The number of consecutive times that attempts to read a manifest or segment must fail before the input is considered unavailable.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The number of seconds between retries when an attempt to read a manifest or segment fails.
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Identifies the source for the SCTE-35 messages that MediaLive will ingest.
	//
	// Messages can be ingested from the content segments (in the stream) or from tags in the playlist (the HLS manifest). MediaLive ignores SCTE-35 information in the source that is not selected.
	Scte35Source *string `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

