package awsmedialivealpha


// Options for a DVB-Sub caption source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   var ocrLanguage OcrLanguage
//
//   dvbSubCaptionSourceOptions := &DvbSubCaptionSourceOptions{
//   	OcrLanguage: ocrLanguage,
//   	Pid: jsii.Number(123),
//   }
//
// Experimental.
type DvbSubCaptionSourceOptions struct {
	// The OCR language to use when converting this image-based source to text.
	// Default: - MediaLive service default.
	//
	// Experimental.
	OcrLanguage OcrLanguage `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// The PID of the source content.
	//
	// Unused for DVB-Sub passthrough.
	// Default: - MediaLive service default.
	//
	// Experimental.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

