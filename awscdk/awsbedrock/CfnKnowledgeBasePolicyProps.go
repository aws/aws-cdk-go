package awsbedrock


// Properties for defining a `CfnKnowledgeBasePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnKnowledgeBasePolicyProps := &CfnKnowledgeBasePolicyProps{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html
//
type CfnKnowledgeBasePolicyProps struct {
	// The unique identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html#cfn-bedrock-knowledgebasepolicy-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The IAM policy document defining access permissions for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html#cfn-bedrock-knowledgebasepolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

