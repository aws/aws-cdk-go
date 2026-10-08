package interfacesawsram


// A reference to a PermissionAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionAssociationReference := &PermissionAssociationReference{
//   	PermissionArn: jsii.String("permissionArn"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
type PermissionAssociationReference struct {
	// The PermissionArn of the PermissionAssociation resource.
	PermissionArn *string `field:"required" json:"permissionArn" yaml:"permissionArn"`
	// The ResourceShareArn of the PermissionAssociation resource.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}

