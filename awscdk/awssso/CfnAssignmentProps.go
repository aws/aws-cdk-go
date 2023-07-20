package awssso


// Properties for defining a `CfnAssignment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssignmentProps := &CfnAssignmentProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	PermissionSetArn: jsii.String("permissionSetArn"),
//   	PrincipalId: jsii.String("principalId"),
//   	PrincipalType: jsii.String("principalType"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html
//
type CfnAssignmentProps struct {
	// The ARN of the IAM Identity Center instance under which the operation will be executed.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The ARN of the permission set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-permissionsetarn
	//
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// An identifier for an object in IAM Identity Center, such as a user or group.
	//
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For more information about PrincipalIds in IAM Identity Center, see the [IAM Identity Center Identity Store API Reference](https://docs.aws.amazon.com//singlesignon/latest/IdentityStoreAPIReference/welcome.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-principalid
	//
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The entity type for which the assignment will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-principaltype
	//
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// TargetID is an AWS account identifier, (For example, 123456789012).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-targetid
	//
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The entity type for which the assignment will be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-assignment.html#cfn-sso-assignment-targettype
	//
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

