package mixinsawsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   audioLanguageConfigurationProperty := &AudioLanguageConfigurationProperty{
//   	GenerativeOutputLanguage: jsii.String("generativeOutputLanguage"),
//   	IdentifyMultipleLanguages: jsii.Boolean(false),
//   	InputLanguages: []*string{
//   		jsii.String("inputLanguages"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_AudioLanguageConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-generativeoutputlanguage
	//
	GenerativeOutputLanguage *string `field:"optional" json:"generativeOutputLanguage" yaml:"generativeOutputLanguage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-identifymultiplelanguages
	//
	IdentifyMultipleLanguages interface{} `field:"optional" json:"identifyMultipleLanguages" yaml:"identifyMultipleLanguages"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-inputlanguages
	//
	InputLanguages *[]*string `field:"optional" json:"inputLanguages" yaml:"inputLanguages"`
}

