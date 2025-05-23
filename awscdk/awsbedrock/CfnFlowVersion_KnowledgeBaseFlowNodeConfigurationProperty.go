package awsbedrock


// Contains configurations for a knowledge base node in a flow.
//
// This node takes a query as the input and returns, as the output, the retrieved responses directly (as an array) or a response generated based on the retrieved responses. For more information, see [Node types in Amazon Bedrock works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   knowledgeBaseFlowNodeConfigurationProperty := &KnowledgeBaseFlowNodeConfigurationProperty{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//
//   	// the properties below are optional
//   	GuardrailConfiguration: &GuardrailConfigurationProperty{
//   		GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   		GuardrailVersion: jsii.String("guardrailVersion"),
//   	},
//   	ModelId: jsii.String("modelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html
//
type CfnFlowVersion_KnowledgeBaseFlowNodeConfigurationProperty struct {
	// The unique identifier of the knowledge base to query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// Contains configurations for a guardrail to apply during query and response generation for the knowledge base in this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-guardrailconfiguration
	//
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) to use to generate a response from the query results. Omit this field if you want to return the retrieved results as an array.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-knowledgebaseflownodeconfiguration.html#cfn-bedrock-flowversion-knowledgebaseflownodeconfiguration-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}

