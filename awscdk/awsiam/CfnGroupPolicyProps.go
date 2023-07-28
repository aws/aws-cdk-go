package awsiam


// Properties for defining a `CfnGroupPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnGroupPolicyProps := &CfnGroupPolicyProps{
//   	GroupName: jsii.String("groupName"),
//   	PolicyName: jsii.String("policyName"),
//
//   	// the properties below are optional
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html
//
type CfnGroupPolicyProps struct {
	// The name of the group to associate the policy with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html#cfn-iam-grouppolicy-groupname
	//
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// The name of the policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html#cfn-iam-grouppolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// The policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-grouppolicy.html#cfn-iam-grouppolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}

