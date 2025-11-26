package previewawsiammixins


// Contains information about an attached policy.
//
// An attached policy is a managed policy that has been attached to a user, group, or role.
//
// For more information about managed policies, see [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *IAM User Guide* .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group-policy.html
//
type CfnGroupPropsMixin_PolicyProperty struct {
	// The policy document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group-policy.html#cfn-iam-group-policy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The friendly name (not ARN) identifying the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group-policy.html#cfn-iam-group-policy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
}

