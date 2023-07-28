package awsiam


// Properties for defining a `CfnRolePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnRolePolicyProps := &CfnRolePolicyProps{
//   	PolicyName: jsii.String("policyName"),
//   	RoleName: jsii.String("roleName"),
//
//   	// the properties below are optional
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html
//
type CfnRolePolicyProps struct {
	// The friendly name (not ARN) identifying the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html#cfn-iam-rolepolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The name of the policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html#cfn-iam-rolepolicy-rolename
	//
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-rolepolicy.html#cfn-iam-rolepolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

