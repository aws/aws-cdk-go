package awssso


// A reference to a PermissionSet resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionSetReference := &PermissionSetReference{
//   	InstanceArn: jsii.String("instanceArn"),
//   	PermissionSetArn: jsii.String("permissionSetArn"),
//   }
//
type PermissionSetReference struct {
	// The InstanceArn of the PermissionSet resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The PermissionSetArn of the PermissionSet resource.
	PermissionSetArn *string `field:"required" json:"permissionSetArn" yaml:"permissionSetArn"`
}

