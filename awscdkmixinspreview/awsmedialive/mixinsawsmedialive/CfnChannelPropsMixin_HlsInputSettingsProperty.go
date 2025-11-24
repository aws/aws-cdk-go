package mixinsawsmedialive


// Information about how to connect to the upstream system.
//
// The parent of this entity is NetworkInputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsInputSettingsProperty := &HlsInputSettingsProperty{
//   	Bandwidth: jsii.Number(123),
//   	BufferSegments: jsii.Number(123),
//   	Retries: jsii.Number(123),
//   	RetryInterval: jsii.Number(123),
//   	Scte35Source: jsii.String("scte35Source"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html
//
type CfnChannelPropsMixin_HlsInputSettingsProperty struct {
	// When specified, the HLS stream with the m3u8 bandwidth that most closely matches this value is chosen.
	//
	// Otherwise, the highest bandwidth stream in the m3u8 is chosen. The bitrate is specified in bits per second, as in an HLS manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html#cfn-medialive-channel-hlsinputsettings-bandwidth
	//
	Bandwidth *float64 `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// When specified, reading of the HLS input begins this many buffer segments from the end (most recently written segment).
	//
	// When not specified, the HLS input begins with the first segment specified in the m3u8.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html#cfn-medialive-channel-hlsinputsettings-buffersegments
	//
	BufferSegments *float64 `field:"optional" json:"bufferSegments" yaml:"bufferSegments"`
	// The number of consecutive times that attempts to read a manifest or segment must fail before the input is considered unavailable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html#cfn-medialive-channel-hlsinputsettings-retries
	//
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The number of seconds between retries when an attempt to read a manifest or segment fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html#cfn-medialive-channel-hlsinputsettings-retryinterval
	//
	RetryInterval *float64 `field:"optional" json:"retryInterval" yaml:"retryInterval"`
	// Identifies the source for the SCTE-35 messages that MediaLive will ingest.
	//
	// Messages can be ingested from the content segments (in the stream) or from tags in the playlist (the HLS manifest). MediaLive ignores SCTE-35 information in the source that is not selected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlsinputsettings.html#cfn-medialive-channel-hlsinputsettings-scte35source
	//
	Scte35Source *string `field:"optional" json:"scte35Source" yaml:"scte35Source"`
}

