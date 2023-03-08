package awsmediapackage


// Limitations for outputs from the endpoint, based on the video bitrate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamSelectionProperty := &streamSelectionProperty{
//   	maxVideoBitsPerSecond: jsii.Number(123),
//   	minVideoBitsPerSecond: jsii.Number(123),
//   	streamOrder: jsii.String("streamOrder"),
//   }
//
type CfnPackagingConfiguration_StreamSelectionProperty struct {
	// The upper limit of the bitrates that this endpoint serves.
	//
	// If the video track exceeds this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 2147483647 bits per second.
	MaxVideoBitsPerSecond *float64 `field:"optional" json:"maxVideoBitsPerSecond" yaml:"maxVideoBitsPerSecond"`
	// The lower limit of the bitrates that this endpoint serves.
	//
	// If the video track is below this threshold, then AWS Elemental MediaPackage excludes it from output. If you don't specify a value, it defaults to 0 bits per second.
	MinVideoBitsPerSecond *float64 `field:"optional" json:"minVideoBitsPerSecond" yaml:"minVideoBitsPerSecond"`
	// Order in which the different video bitrates are presented to the player.
	//
	// Valid values: `ORIGINAL` , `VIDEO_BITRATE_ASCENDING` , `VIDEO_BITRATE_DESCENDING` .
	StreamOrder *string `field:"optional" json:"streamOrder" yaml:"streamOrder"`
}

