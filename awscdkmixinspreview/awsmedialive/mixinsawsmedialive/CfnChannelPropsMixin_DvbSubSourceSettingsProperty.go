package mixinsawsmedialive


// Information about the DVB Sub captions to extract from the input.
//
// The parent of this entity is CaptionSelectorSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dvbSubSourceSettingsProperty := &DvbSubSourceSettingsProperty{
//   	OcrLanguage: jsii.String("ocrLanguage"),
//   	Pid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubsourcesettings.html
//
type CfnChannelPropsMixin_DvbSubSourceSettingsProperty struct {
	// If you will configure a WebVTT caption description that references this caption selector, use this field to provide the language to consider when translating the image-based source to text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubsourcesettings.html#cfn-medialive-channel-dvbsubsourcesettings-ocrlanguage
	//
	OcrLanguage *string `field:"optional" json:"ocrLanguage" yaml:"ocrLanguage"`
	// When using DVB-Sub with burn-in or SMPTE-TT, use this PID for the source content.
	//
	// It is unused for DVB-Sub passthrough. All DVB-Sub content is passed through, regardless of selectors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsubsourcesettings.html#cfn-medialive-channel-dvbsubsourcesettings-pid
	//
	Pid *float64 `field:"optional" json:"pid" yaml:"pid"`
}

