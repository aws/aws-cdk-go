package previewawsbedrockmixins


// Contains configurations for a text prompt template.
//
// To include a variable, enclose a word in double curly braces as in `{{variable}}` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   textPromptTemplateConfigurationProperty := &TextPromptTemplateConfigurationProperty{
//   	CachePoint: &CachePointBlockProperty{
//   		Type: jsii.String("type"),
//   	},
//   	InputVariables: []interface{}{
//   		&PromptInputVariableProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Text: jsii.String("text"),
//   	TextS3Location: &TextS3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html
//
type CfnPromptPropsMixin_TextPromptTemplateConfigurationProperty struct {
	// A cache checkpoint within a template configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-cachepoint
	//
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// An array of the variables in the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-inputvariables
	//
	InputVariables interface{} `field:"optional" json:"inputVariables" yaml:"inputVariables"`
	// The message for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The Amazon S3 location of the prompt text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-textprompttemplateconfiguration.html#cfn-bedrock-prompt-textprompttemplateconfiguration-texts3location
	//
	TextS3Location interface{} `field:"optional" json:"textS3Location" yaml:"textS3Location"`
}

