package awslogs


// Properties for defining a `CfnAccountPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountPolicyProps := &CfnAccountPolicyProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html
//
type CfnAccountPolicyProps struct {
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per PolicyType.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policydocument
	//
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The name of the account policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Type of the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policytype
	//
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Scope for policy application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

