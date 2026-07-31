package awsmediaconnectalpha


// Options for Media Stream Source Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   var encoding Encoding
//   var mediaStream MediaStream
//
//   mediaStreamSourceConfigurationCdi := &MediaStreamSourceConfigurationCdi{
//   	Encoding: encoding,
//   	MediaStream: mediaStream,
//   }
//
// Experimental.
type MediaStreamSourceConfigurationCdi struct {
	// The format that was used to encode the data.
	//
	// For ancillary data streams, set the encoding name to smpte291.
	// If the encoding name is smpte291, set the color space to one of the following:
	// BT601, BT709, BT2020, or BT2100. For all other ancillary data streams, set the color space to SDR-NOCOLOR.
	// Experimental.
	Encoding Encoding `field:"required" json:"encoding" yaml:"encoding"`
	// The name of the media stream.
	// Experimental.
	MediaStream MediaStream `field:"required" json:"mediaStream" yaml:"mediaStream"`
}

