package awsbedrock


// Settings for an Amazon Kendra knowledge base.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kendraKnowledgeBaseConfigurationProperty := &KendraKnowledgeBaseConfigurationProperty{
//   	KendraIndexArn: jsii.String("kendraIndexArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-kendraknowledgebaseconfiguration.html
//
type CfnKnowledgeBase_KendraKnowledgeBaseConfigurationProperty struct {
	// The ARN of the Amazon Kendra index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-knowledgebase-kendraknowledgebaseconfiguration.html#cfn-bedrock-knowledgebase-kendraknowledgebaseconfiguration-kendraindexarn
	//
	KendraIndexArn *string `field:"required" json:"kendraIndexArn" yaml:"kendraIndexArn"`
}

