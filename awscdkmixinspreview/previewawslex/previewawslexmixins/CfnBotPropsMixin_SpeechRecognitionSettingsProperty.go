package previewawslexmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   speechRecognitionSettingsProperty := &SpeechRecognitionSettingsProperty{
//   	SpeechModelConfig: &SpeechModelConfigProperty{
//   		DeepgramConfig: &DeepgramSpeechModelConfigProperty{
//   			ApiTokenSecretArn: jsii.String("apiTokenSecretArn"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   	},
//   	SpeechModelPreference: jsii.String("speechModelPreference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechrecognitionsettings.html
//
type CfnBotPropsMixin_SpeechRecognitionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechrecognitionsettings.html#cfn-lex-bot-speechrecognitionsettings-speechmodelconfig
	//
	SpeechModelConfig interface{} `field:"optional" json:"speechModelConfig" yaml:"speechModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechrecognitionsettings.html#cfn-lex-bot-speechrecognitionsettings-speechmodelpreference
	//
	SpeechModelPreference *string `field:"optional" json:"speechModelPreference" yaml:"speechModelPreference"`
}

