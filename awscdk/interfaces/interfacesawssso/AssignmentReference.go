package interfacesawssso


// A reference to a Assignment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assignmentReference := &AssignmentReference{
//   	InstanceArn: jsii.String("instanceArn"),
//   	PermissionSetArn: jsii.String("permissionSetArn"),
//   	PrincipalId: jsii.String("principalId"),
//   	PrincipalType: jsii.String("principalType"),
//   	TargetId: jsii.String("targetId"),
//   	TargetType: jsii.String("targetType"),
//   }
//
type AssignmentReference struct {
	// The InstanceArn of the Assignment resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The PermissionSetArn of the Assignment resource.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
	// The PrincipalId of the Assignment resource.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// The PrincipalType of the Assignment resource.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The TargetId of the Assignment resource.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
	// The TargetType of the Assignment resource.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
}

