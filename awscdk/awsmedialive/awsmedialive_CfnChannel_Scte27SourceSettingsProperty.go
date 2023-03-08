package awsmedialive


// Information about the SCTE-27 captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scte27SourceSettingsProperty := &scte27SourceSettingsProperty{
//   	ocrLanguage: jsii.String("ocrLanguage"),
//   	pid: jsii.Number(123),
//   }
//
type CfnChannel_Scte27SourceSettingsProperty struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// The PID field is used in conjunction with the captions selector languageCode field as follows: Specify PID and Language: Extracts captions from that PID;
	//
	// the language is "informational." Specify PID and omit Language: Extracts the specified PID. Omit PID and specify Language: Extracts the specified language, whichever PID that happens to be. Omit PID and omit Language: Valid only if source is DVB-Sub that is being passed through; all languages are passed through.
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

