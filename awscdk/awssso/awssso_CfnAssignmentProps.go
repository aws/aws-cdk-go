package awssso


// Properties for defining a `CfnAssignment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssignmentProps := &cfnAssignmentProps{
//   	instanceArn: jsii.String("instanceArn"),
//   	permissionSetArn: jsii.String("permissionSetArn"),
//   	principalId: jsii.String("principalId"),
//   	principalType: jsii.String("principalType"),
//   	targetId: jsii.String("targetId"),
//   	targetType: jsii.String("targetType"),
//   }
//
type CfnAssignmentProps struct {
	// The ARN of the SSO instance under which the operation will be executed.
	//
	// For more information about ARNs, see [Amazon Resource Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com//general/latest/gr/aws-arns-and-namespaces.html) in the *AWS General Reference* .
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The ARN of the permission set.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// An identifier for an object in AWS SSO , such as a user or group.
	//
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For more information about PrincipalIds in AWS SSO , see the [AWS SSO Identity Store API Reference](https://docs.aws.amazon.com//singlesignon/latest/IdentityStoreAPIReference/welcome.html) .
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The entity type for which the assignment will be created.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// TargetID is an AWS account identifier, typically a 10-12 digit string (For example, 123456789012).
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The entity type for which the assignment will be created.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

