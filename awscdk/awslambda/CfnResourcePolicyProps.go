package awslambda


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-resourcepolicy.html
//
type CfnResourcePolicyProps struct {
	// The policy document you want to add to your LAM resource.
	//
	// This is formatted as a JSON string.
	//  For more information, see [Working with resource-based policies in](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) in the *Developer Guide*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-resourcepolicy.html#cfn-lambda-resourcepolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Name (ARN) of the LAM resource you want to add the policy to.
	//
	// For a function, you can use a qualified or an unqualified ARN. The value must be a complete ARN, and the operation does not accept wildcard characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-resourcepolicy.html#cfn-lambda-resourcepolicy-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

