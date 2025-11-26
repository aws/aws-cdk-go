package previewawslexmixins


// Provides settings that enable advanced recognition settings for slot values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedRecognitionSettingProperty := &AdvancedRecognitionSettingProperty{
//   	AudioRecognitionStrategy: jsii.String("audioRecognitionStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-advancedrecognitionsetting.html
//
type CfnBotPropsMixin_AdvancedRecognitionSettingProperty struct {
	// Enables using the slot values as a custom vocabulary for recognizing user utterances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-advancedrecognitionsetting.html#cfn-lex-bot-advancedrecognitionsetting-audiorecognitionstrategy
	//
	AudioRecognitionStrategy *string `field:"optional" json:"audioRecognitionStrategy" yaml:"audioRecognitionStrategy"`
}

