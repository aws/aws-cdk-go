package awsmedialivealpha


// Options for an SCTE-27 caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var ocrLanguage OcrLanguage
//
//   scte27CaptionSourceOptions := &Scte27CaptionSourceOptions{
//   	OcrLanguage: ocrLanguage,
//   	Pid: jsii.Number(123),
//   }
//
// Experimental.
type Scte27CaptionSourceOptions struct {
	// The OCR language to use when converting this image-based source to text.
	// Default: - MediaLive service default.
	//
	// Experimental.
	OcrLanguage OcrLanguage `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// The PID to extract captions from.
	//
	// See the MediaLive docs for how PID and `languageCode`
	// interact.
	// Default: - MediaLive service default.
	//
	// Experimental.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

