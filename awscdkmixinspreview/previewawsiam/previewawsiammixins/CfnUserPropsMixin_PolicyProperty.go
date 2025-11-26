package previewawsiammixins


// Contains information about an attached policy.
//
// An attached policy is a managed policy that has been attached to a user, group, or role.
//
// For more information about managed policies, refer to [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   policyProperty := &PolicyProperty{
//   	PolicyDocument: policyDocument,
//   	PolicyName: jsii.String("policyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-policy.html
//
type CfnUserPropsMixin_PolicyProperty struct {
	// The entire contents of the policy that defines permissions.
	//
	// For more information, see [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-policy.html#cfn-iam-user-policy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-policy.html#cfn-iam-user-policy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

