package awsbedrock


// Defines a custom prompt template for orchestrating the retrieval and generation process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBasePromptTemplateProperty := &KnowledgeBasePromptTemplateProperty{
//   	TextPromptTemplate: jsii.String("textPromptTemplate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseprompttemplate.html
//
type CfnFlowVersion_KnowledgeBasePromptTemplateProperty struct {
	// The text of the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseprompttemplate.html#cfn-bedrock-flowversion-knowledgebaseprompttemplate-textprompttemplate
	//
	TextPromptTemplate *string `field:"required" json:"textPromptTemplate" yaml:"textPromptTemplate"`
}

