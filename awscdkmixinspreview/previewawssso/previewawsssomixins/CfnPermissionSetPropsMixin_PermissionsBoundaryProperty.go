package previewawsssomixins


// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary.
//
// Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
//
// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   permissionsBoundaryProperty := &PermissionsBoundaryProperty{
//   	CustomerManagedPolicyReference: &CustomerManagedPolicyReferenceProperty{
//   		Name: jsii.String("name"),
//   		Path: jsii.String("path"),
//   	},
//   	ManagedPolicyArn: jsii.String("managedPolicyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-permissionset-permissionsboundary.html
//
type CfnPermissionSetPropsMixin_PermissionsBoundaryProperty struct {
	// Specifies the name and path of a customer managed policy.
	//
	// You must have an IAM policy that matches the name and path in each AWS account where you want to deploy your permission set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-permissionset-permissionsboundary.html#cfn-sso-permissionset-permissionsboundary-customermanagedpolicyreference
	//
	CustomerManagedPolicyReference interface{} `field:"optional" json:"customerManagedPolicyReference" yaml:"customerManagedPolicyReference"`
	// The AWS managed policy ARN that you want to attach to a permission set as a permissions boundary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sso-permissionset-permissionsboundary.html#cfn-sso-permissionset-permissionsboundary-managedpolicyarn
	//
	ManagedPolicyArn *string `field:"optional" json:"managedPolicyArn" yaml:"managedPolicyArn"`
}

