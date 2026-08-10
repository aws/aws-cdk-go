package awsbedrock


// Properties for CfnKnowledgeBasePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var policyDocument interface{}
//
//   cfnKnowledgeBasePolicyMixinProps := &CfnKnowledgeBasePolicyMixinProps{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html
//
type CfnKnowledgeBasePolicyMixinProps struct {
	// The unique identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html#cfn-bedrock-knowledgebasepolicy-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The IAM policy document defining access permissions for the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-knowledgebasepolicy.html#cfn-bedrock-knowledgebasepolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

