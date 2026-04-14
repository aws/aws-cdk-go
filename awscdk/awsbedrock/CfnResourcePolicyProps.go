package awsbedrock


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnResourcePolicyProps := &CfnResourcePolicyProps{
//   	PolicyDocument: policyDocument,
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-resourcepolicy.html
//
type CfnResourcePolicyProps struct {
	// The IAM policy document defining access permissions for the guardrail and guardrail profile resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-resourcepolicy.html#cfn-bedrock-resourcepolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The ARN of the Bedrock Guardrail or Guardrail Profile resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-resourcepolicy.html#cfn-bedrock-resourcepolicy-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

