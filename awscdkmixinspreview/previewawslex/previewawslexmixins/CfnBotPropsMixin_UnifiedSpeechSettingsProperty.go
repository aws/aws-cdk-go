package previewawslexmixins


// Unified configuration settings that combine speech recognition and synthesis capabilities.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   unifiedSpeechSettingsProperty := &UnifiedSpeechSettingsProperty{
//   	SpeechFoundationModel: &SpeechFoundationModelProperty{
//   		ModelArn: jsii.String("modelArn"),
//   		VoiceId: jsii.String("voiceId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-unifiedspeechsettings.html
//
type CfnBotPropsMixin_UnifiedSpeechSettingsProperty struct {
	// The foundation model configuration to use for unified speech processing capabilities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-unifiedspeechsettings.html#cfn-lex-bot-unifiedspeechsettings-speechfoundationmodel
	//
	SpeechFoundationModel interface{} `field:"optional" json:"speechFoundationModel" yaml:"speechFoundationModel"`
}

