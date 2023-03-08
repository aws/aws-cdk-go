package awsmedialive


// Information about the DVB Sub captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSubSourceSettingsProperty := &dvbSubSourceSettingsProperty{
//   	ocrLanguage: jsii.String("ocrLanguage"),
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_DvbSubSourceSettingsProperty struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// When using DVB-Sub with burn-in or SMPTE-TT, use this PID for the source content.
	//
	// It is unused for DVB-Sub passthrough. All DVB-Sub content is passed through, regardless of selectors.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

