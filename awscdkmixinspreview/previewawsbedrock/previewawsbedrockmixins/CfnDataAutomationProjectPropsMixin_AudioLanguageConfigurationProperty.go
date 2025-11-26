package previewawsbedrockmixins


// This allows you to set the input and output language of your audio.
//
// The input language can be set to any of the languages supported by Bedrock Data Automation. The output can either be set to english or whatever the dominant language is of the audio, determined by the language spoken for the most seconds.
//
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
	// The output language of your processing results.
	//
	// This can either be set to `EN` (English) or `DEFAULT` which will output the results in the dominant language of the audio. The dominant language is determined as the language in the audio, spoken the longest in the input audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-generativeoutputlanguage
	//
	GenerativeOutputLanguage *string `field:"optional" json:"generativeOutputLanguage" yaml:"generativeOutputLanguage"`
	// The toggle determining if you want to detect multiple languages from your audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-identifymultiplelanguages
	//
	IdentifyMultipleLanguages interface{} `field:"optional" json:"identifyMultipleLanguages" yaml:"identifyMultipleLanguages"`
	// The input language of your audio.
	//
	// This can be set to any of the currently supported languages via the language codes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiolanguageconfiguration.html#cfn-bedrock-dataautomationproject-audiolanguageconfiguration-inputlanguages
	//
	InputLanguages *[]*string `field:"optional" json:"inputLanguages" yaml:"inputLanguages"`
}

