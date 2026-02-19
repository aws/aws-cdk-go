package previewawslexmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   speechModelConfigProperty := &SpeechModelConfigProperty{
//   	DeepgramConfig: &DeepgramSpeechModelConfigProperty{
//   		ApiTokenSecretArn: jsii.String("apiTokenSecretArn"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechmodelconfig.html
//
type CfnBotPropsMixin_SpeechModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-speechmodelconfig.html#cfn-lex-bot-speechmodelconfig-deepgramconfig
	//
	DeepgramConfig interface{} `field:"optional" json:"deepgramConfig" yaml:"deepgramConfig"`
}

