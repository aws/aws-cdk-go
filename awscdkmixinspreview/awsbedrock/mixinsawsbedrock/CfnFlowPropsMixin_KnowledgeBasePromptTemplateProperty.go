package mixinsawsbedrock


// Defines a custom prompt template for orchestrating the retrieval and generation process.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   knowledgeBasePromptTemplateProperty := &KnowledgeBasePromptTemplateProperty{
//   	TextPromptTemplate: jsii.String("textPromptTemplate"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseprompttemplate.html
//
type CfnFlowPropsMixin_KnowledgeBasePromptTemplateProperty struct {
	// The text of the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-knowledgebaseprompttemplate.html#cfn-bedrock-flow-knowledgebaseprompttemplate-textprompttemplate
	//
	TextPromptTemplate *string `field:"optional" json:"textPromptTemplate" yaml:"textPromptTemplate"`
}

