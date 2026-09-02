package awsmedialivealpha


// A reference to an SDP file that describes a SMPTE 2110 stream to ingest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   smpte2110SdpLocation := &Smpte2110SdpLocation{
//   	SdpUrl: jsii.String("sdpUrl"),
//
//   	// the properties below are optional
//   	MediaIndex: jsii.Number(123),
//   }
//
// Experimental.
type Smpte2110SdpLocation struct {
	// The URL of the SDP file.
	// Experimental.
	SdpUrl *string `field:"required" json:"sdpUrl" yaml:"sdpUrl"`
	// The index of the media stream within the SDP to ingest.
	//
	// Use when the SDP describes
	// more than one stream of that media type.
	// Default: - service default.
	//
	// Experimental.
	MediaIndex *float64 `field:"optional" json:"mediaIndex" yaml:"mediaIndex"`
}

