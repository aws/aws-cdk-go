package awsmedialivealpha


// Properties for a SMPTE 2110 receiver group input. Requires `anywhereSettings` on the channel.
//
// You specify the SDP files that describe the streams to ingest — one video SDP, and any
// number of audio and ancillary SDPs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   smpte2110InputProps := &Smpte2110InputProps{
//   	AncillarySdps: []Smpte2110SdpLocation{
//   		&Smpte2110SdpLocation{
//   			SdpUrl: jsii.String("sdpUrl"),
//
//   			// the properties below are optional
//   			MediaIndex: jsii.Number(123),
//   		},
//   	},
//   	AudioSdps: []Smpte2110SdpLocation{
//   		&Smpte2110SdpLocation{
//   			SdpUrl: jsii.String("sdpUrl"),
//
//   			// the properties below are optional
//   			MediaIndex: jsii.Number(123),
//   		},
//   	},
//   	VideoSdp: &Smpte2110SdpLocation{
//   		SdpUrl: jsii.String("sdpUrl"),
//
//   		// the properties below are optional
//   		MediaIndex: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type Smpte2110InputProps struct {
	// The SDPs describing the ancillary data streams (SCTE-35 or captions).
	//
	// Up to 50.
	// Default: - no ancillary data streams.
	//
	// Experimental.
	AncillarySdps *[]*Smpte2110SdpLocation `field:"optional" json:"ancillarySdps" yaml:"ancillarySdps"`
	// The SDPs describing the audio streams.
	//
	// Up to 50.
	// Default: - no audio streams.
	//
	// Experimental.
	AudioSdps *[]*Smpte2110SdpLocation `field:"optional" json:"audioSdps" yaml:"audioSdps"`
	// The SDP describing the video stream.
	// Default: - no video stream.
	//
	// Experimental.
	VideoSdp *Smpte2110SdpLocation `field:"optional" json:"videoSdp" yaml:"videoSdp"`
}

