package previewawsssomixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPermissionSetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var inlinePolicy interface{}
//
//   cfnPermissionSetMixinProps := &CfnPermissionSetMixinProps{
//   	CustomerManagedPolicyReferences: []interface{}{
//   		&CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	InlinePolicy: inlinePolicy,
//   	InstanceArn: jsii.String("instanceArn"),
//   	ManagedPolicies: []*string{
//   		jsii.String("managedPolicies"),
//   	},
//   	Name: jsii.String("name"),
//   	PermissionsBoundary: &PermissionsBoundaryProperty{
//   		CustomerManagedPolicyReference: &CustomerManagedPolicyReferenceProperty{
//   			Name: jsii.String("name"),
//   			Path: jsii.String("path"),
//   		},
//   		ManagedPolicyArn: jsii.String("managedPolicyArn"),
//   	},
//   	RelayStateType: jsii.String("relayStateType"),
//   	SessionDuration: jsii.String("sessionDuration"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html
//
type CfnPermissionSetMixinProps struct {
	// Specifies the names and paths of the customer managed policies that you have attached to your permission set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-customermanagedpolicyreferences
	//
	CustomerManagedPolicyReferences interface{} `field:"optional" json:"customerManagedPolicyReferences" yaml:"customerManagedPolicyReferences"`
	// The description of the `PermissionSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The inline policy that is attached to the permission set.
	//
	// > For `Length Constraints` , if a valid ARN is provided for a permission set, it is possible for an empty inline policy to be returned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-inlinepolicy
	//
	InlinePolicy interface{} `field:"optional" json:"inlinePolicy" yaml:"inlinePolicy"`
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// A structure that stores a list of managed policy ARNs that describe the associated AWS managed policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-managedpolicies
	//
	ManagedPolicies *[]*string `field:"optional" json:"managedPolicies" yaml:"managedPolicies"`
	// The name of the permission set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the configuration of the AWS managed or customer managed policy that you want to set as a permissions boundary.
	//
	// Specify either `CustomerManagedPolicyReference` to use the name and path of a customer managed policy, or `ManagedPolicyArn` to use the ARN of an AWS managed policy. A permissions boundary represents the maximum permissions that any policy can grant your role. For more information, see [Permissions boundaries for IAM entities](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html) in the *IAM User Guide* .
	//
	// > Policies used as permissions boundaries don't provide permissions. You must also attach an IAM policy to the role. To learn how the effective permissions for a role are evaluated, see [IAM JSON policy evaluation logic](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html) in the *IAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-permissionsboundary
	//
	PermissionsBoundary interface{} `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// Used to redirect users within the application during the federation authentication process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-relaystatetype
	//
	RelayStateType *string `field:"optional" json:"relayStateType" yaml:"relayStateType"`
	// The length of time that the application user sessions are valid for in the ISO-8601 standard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-sessionduration
	//
	SessionDuration *string `field:"optional" json:"sessionDuration" yaml:"sessionDuration"`
	// The tags to attach to the new `PermissionSet` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

